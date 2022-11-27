import React, { useState } from 'react';
import * as Icon from 'react-icons/fa';
import { Alert, Button, Form, Modal } from 'react-bootstrap';
import { API } from '../../Config/Api';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';

function AddProduct({ showProduct, setShowProduct, refetchProduct }) {
    const handleClose = () => setShowProduct(false)
    const [preview, setPreview] = useState(null);
    const [message, setMessage] = useState(null)
    const navigate = useNavigate()

    const [product, setProduct] = useState({
        image: "",
        nama: "",
        harga_beli: 0,
        harga_jual: 0,
        stok: 0,
    })

    const handleOnChange = (e) => {
        setProduct({
            ...product,
            [e.target.name]:
                e.target.type === "file" ? e.target.files : e.target.value,
        })
        if (e.target.type === "file") {
            let url = URL.createObjectURL(e.target.files[0]);
            setPreview(url)
        }
    }

    const handleOnSubmit = useMutation(async (e) => {
        try {
            e.preventDefault();

            const formData = new FormData();
            formData.set("image", product.image[0], product.image[0].name);
            formData.set("nama", product.nama);
            formData.set("harga_beli", product.harga_beli);
            formData.set("harga_jual", product.harga_jual);
            formData.set("stok", product.stok);

            const data = await API.post("/Product", formData, {
                headers: {
                    Authorization: `Bearer ${localStorage.token}`,
                },
            });
            if (data.data.code === 200) {
                const alert = (
                    <Alert variant="success" className="py-1 fw-bold">
                        Produk berhasil ditambah
                    </Alert>
                )
                refetchProduct()
                setMessage(alert);
                setPreview(null)
                setProduct({
                    image: "",
                    nama: "",
                    harga_beli: 0,
                    harga_jual: 0,
                    stok: 0,
                })

                const timer = setTimeout(navigates, 1000);

                function navigates() {
                    setMessage(null)
                    setShowProduct(false)
                    navigate("/Dashboard");
                }


            }
        } catch (error) {

            console.log(error.data.data.message);
            const alert = (
                <Alert variant="danger" className="py-1">
                    {error.data.data.message}
                </Alert>
            )
            setMessage(alert);
        }
    })


    return (
        <Modal show={showProduct} onHide={handleClose} size="md" centered>
            <Modal.Header closeButton>
                <Modal.Title className='modal-title fs-1 fw-bold' style={{ color: "#ff4b00" }}>Add Produk</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {message && message}
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
                            value={product.nama}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicHargaBeli">
                        <Form.Label>Harga Beli</Form.Label>
                        <Form.Control
                            type="number"
                            name="harga_beli"
                            placeholder="Harga beli"
                            onChange={handleOnChange}
                            value={product.harga_beli}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicHargaJual">
                        <Form.Label>Harga Jual</Form.Label>
                        <Form.Control
                            type="number"
                            name="harga_jual"
                            placeholder="Harga jual"
                            onChange={handleOnChange}
                            value={product.harga_jual}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicStok">
                        <Form.Label>Stok</Form.Label>
                        <Form.Control
                            type="number"
                            name="stok"
                            placeholder="Stok"
                            onChange={handleOnChange}
                            value={product.stok}
                        />
                    </Form.Group>
                    <Button type='submit' className='btn text-white fw-bold link w-100 border-0' style={{ background: "#ff4b00" }}>Save</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}

export default AddProduct;