import React, { useContext, useEffect, useState } from 'react'
import { Form, Button, Table } from 'react-bootstrap';
import * as Icon from 'react-icons/fa';
import { useMutation, useQuery } from 'react-query';
import { API } from '../Config/Api';
import { UserContext } from '../Contex/User-context';
import AddProduct from './Modal/Add-product';
import Deleteproduct from './Modal/Delete-product';
import EditProduct from './Modal/Edit-Product';

function Product() {
    const [showAddProduct, setShowAddProduct] = useState(false)
    const [showEditProduct, setShowEditProduct] = useState(false)
    const [state, dispatch] = useContext(UserContext)
    const [idDelete, setIdDelete] = useState(null);
    const [idEdit, setIdEdit] = useState(null);
    const [confirmDelete, setConfirmDelete] = useState(null);
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const userId = state.user.id

    let { data: dataProduct, refetch } = useQuery("dataProductCache", async () => {
        const response = await API.get("/ProductByUser/" + userId)
        return response.data.data
    })

    const handleEdit = (id) => {
        setIdEdit(id);
        setShowEditProduct(true);
    };

    const handleDelete = (id) => {
        setIdDelete(id);
        handleShow();
    };

    const deleteById = useMutation(async (id) => {
        try {
            await API.delete(`/Product/${id}`);
            refetch();
        } catch (error) {
            console.log(error);
        }
    });

    useEffect(() => {
        if (confirmDelete) {
            handleClose();
            deleteById.mutate(idDelete);
            setConfirmDelete(null);
        }
    }, [confirmDelete]);

    const [dataEdit, setDataEdit] = useState()

    const getEdit = async (id) => {
        try {
            const response = await API.get("/Product/" + id)
            setDataEdit(response.data.data)
        } catch (error) {
            console.log(error.response);
        }
    }
    // console.log(dataEdit);



    return (
        <div className='container'>
            <div className="Product mt-3">
                <p className="fs-3 fw-bold">List Produk</p>
                <hr />
                <div className="top d-md-flex justify-content-between">
                    <Form>
                        <Form.Group className="mb-3" controlId="formBasicSearch">
                            <Form.Control
                                type="text"
                                name="search"
                                placeholder="Search ..."
                            // onChange={handleOnChange}
                            // value={formLogin.username}
                            />
                        </Form.Group>
                    </Form>
                    <div className="tambah">
                        <Button onClick={() => setShowAddProduct(true)}><Icon.FaPlusCircle /> Tambah</Button>
                    </div>
                </div>
            </div>

            <Table striped bordered hover variant='primary' >
                <thead>
                    <tr className='text-center'>
                        <th>No</th>
                        <th>Foto Barang</th>
                        <th>Nama Barang</th>
                        <th>Harga Beli</th>
                        <th>Harga Jual</th>
                        <th>Stok</th>
                        <th colSpan={2}>Opsi</th>
                    </tr>
                </thead>
                <tbody>
                    {dataProduct?.map((item, index) => (
                        <tr key={index}>
                            <td>{index + 1}</td>
                            <td><img
                                src={item?.image}
                                style={{
                                    width: "80px",
                                    height: "80px",
                                    objectFit: "cover",
                                    borderRadius: "2px",
                                }}
                                alt={item?.nama}
                            /></td>
                            <td>{item?.nama}</td>
                            <td>{item?.harga_beli}</td>
                            <td>{item?.harga_jual}</td>
                            <td>{item?.stok}</td>
                            <td className='text-center'
                                onClick={() => {
                                    handleEdit(item.id);
                                    getEdit(item.id)
                                }} >
                                <Icon.FaRegEdit />
                            </td>
                            <td className='text-center'
                                onClick={() => {
                                    handleDelete(item.id);
                                }} >
                                <Icon.FaRegTrashAlt />
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>

            <AddProduct showProduct={showAddProduct} setShowProduct={setShowAddProduct} refetchProduct={refetch} />
            <EditProduct showEdit={showEditProduct} setShowEdit={setShowEditProduct} idEdit={idEdit} dataEdit={dataEdit} />
            <Deleteproduct setConfirmDelete={setConfirmDelete} show={show} handleClose={handleClose} />
        </div>
    )
}

export default Product;