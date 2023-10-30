import React, { useState, useEffect, useCallback } from 'react'
import { Button,Form,Input,Row,Col,Divider } from 'antd';
import { update } from './service';
import {useNavigate,useLocation } from "react-router-dom"
import { fetchTheData } from './service';


export function EditDocuments(){

    const location=useLocation();
    const [isCancelModalOpen, setIsCancelModalOpen] = useState(false);
    const[values,setValues]=useState({title:'',format:'',author:'',owner:'',validity:''})
    var data
    data=location.state


    useEffect(() => {
        fetch(data);
    },[isCancelModalOpen] );


    const fetch = (data) => {
    const resp = fetchTheData(data);
    resp
      .then((res) => {
        console.log(res.data)
        const updatedValues = { ...values }; 
        updatedValues.id=res.data.id
        updatedValues.title = res.data.title;
        updatedValues.format = res.data.format;
        updatedValues.author = res.data.author;
        updatedValues.owner = res.data.owner;
        updatedValues.validity = res.data.validity;
  
        setValues(updatedValues); 
      })
      .catch((err) => console.log(err));
  };
  

  
    const [form] = Form.useForm();
    const navigate=useNavigate()
    const [errorName,setErrorName] = useState('')

      
      const set=(event)=>{
        setValues({...values,[event.target.name]:event.target.value})
      }
      

      const handleSubmit=async (e)=>{
        let index=values.id
        console.log(data)

    if(!values.title || !values.format || !values.author || !values.owner || !values.validity ){
        alert("please complete the details")
      }
      else{
        e.preventDefault()
        const resp = update(values);
        resp
        .then((res) => {
            navigate('/')
        })
        .catch((err) =>{ 
            console.log(err)
            setErrorName(err.response.data)
        });  

        }
    }


    return(
        <>
              <div className='leader'>
      <section className='container'>
        <header className='header'>Edit Documents</header>
      <Form className='form'
        labelCol={{
          span: 4,
        }}
        wrapperCol={{
          span: 8,
        }}
        layout="horizontal"
      >
        <div className='inputbox'>
        <Form.Item label="Title">
          <Input name='title' value={values.title} onChange={set} disabled={true}/>
        </Form.Item>
        <Form.Item label="Format">
          <Input name='format' value={values.format} onChange={set}/>
        </Form.Item>
        <Form.Item label="Author">
          <Input name='author' value={values.author} onChange={set} disabled={true}/>
        </Form.Item>
        <Form.Item label="Owner">
          <Input  name='owner' value={values.owner} onChange={set}/>
        </Form.Item>
        <Form.Item label="Validity">
          <Input  name='validity' value={values.validity} onChange={set}/>
        </Form.Item>

        <br/>
        <p style={{color:'red'}}>{errorName}</p>
        <Button type="primary" className='submit' onClick={handleSubmit}>Update</Button>
        </div>
      </Form>
      </section>
      </div>
        </>
    )
}