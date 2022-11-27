import React, { useEffect, useState } from 'react'
import { Button, Form, Modal } from 'react-bootstrap'
import * as Icon from 'react-icons/fa'
import { useMutation, useQuery } from 'react-query'
import { API } from '../../Config/Api'

function EditProduct({ showEdit, setShowEdit, idEdit, dataEdit }) {
    const handleClose = () => setShowEdit(false)
    const [preview, setPreview] = useState(dataEdit?.image);
    // const [dataEdits, setDataEdit] = useState(dataEdit)
    // console.log('ini data', dataEdits);
    const setDataEdit = useState()

    console.log(dataEdit?.nama);



    const handleOnChange = (e) => {
        // setDataEdit({
        //     ...dataEdit,
        //     [e.target.name]:
        //         e.target.type === "file" ? e.target.files : e.target.value,
        // })
        // if (e.target.type === "file") {
        //     let url = URL.createObjectURL(e.target.files[0]);
        //     setPreview(url)
        // }

        // console.log(dataEdit);
    }


    const handleOnSubmit = useMutation(async (e) => {
        alert("test")
    })


    return (
        <Modal show={showEdit} onHide={handleClose} size="md" centered>
            <Modal.Header closeButton>
                <Modal.Title className='modal-title fs-1 fw-bold' style={{ color: "#ff4b00" }}>Edit Produk</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {/* {message && message} */}
                <Form onSubmit={(e) => handleOnSubmit.mutate(e)}>
                    <Form.Group className='w-100 mb-3' controlId="formBasicimage">
                        <Form.Label className="btn text-white" style={{
                            backgroundColor: "#ff4b00"
                        }}>
                            Upload image &nbsp; <Icon.FaImage />
                        </Form.Label>
                        <Form.Control
                            type="file"
                            name="image"
                            onChange={handleOnChange}
                            hidden
                        />
                    </Form.Group>
                    {preview && (
                        <div>
                            <img className='rounded'
                                src={preview}
                                style={{
                                    maxWidth: "150px",
                                    maxHeight: "150px",
                                    objectFit: "cover",
                                }}
                                alt={preview}
                            />
                        </div>
                    )}
                    <Form.Group className="mb-3" controlId="formBasicNama">
                        <Form.Label>Nama Barang</Form.Label>
                        <Form.Control
                            type="text"
                            name="nama"
                            placeholder="Nama Barang"
                            onChange={handleOnChange}
                            value={dataEdit?.nama}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicHargaBeli">
                        <Form.Label>Harga Beli</Form.Label>
                        <Form.Control
                            type="number"
                            name="harga_beli"
                            placeholder="Harga beli"
                            onChange={handleOnChange}
                            value={dataEdit?.harga_beli}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicHargaJual">
                        <Form.Label>Harga Jual</Form.Label>
                        <Form.Control
                            type="number"
                            name="harga_jual"
                            placeholder="Harga jual"
                            onChange={handleOnChange}
                            value={dataEdit?.harga_jual}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicStok">
                        <Form.Label>Stok</Form.Label>
                        <Form.Control
                            type="number"
                            name="stok"
                            placeholder="Stok"
                            onChange={handleOnChange}
                            value={dataEdit?.stok}
                        />
                    </Form.Group>
                    <Button type='submit' className='btn text-white fw-bold link w-100 border-0' style={{ background: "#ff4b00" }}>Save</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}

export default EditProduct;