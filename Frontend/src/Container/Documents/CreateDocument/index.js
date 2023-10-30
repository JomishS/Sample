import React, { useState } from 'react';
import './AddDocument.scss'
import axios from 'axios';
import {Form,Input,Button} from 'antd';
import { Link,useNavigate } from "react-router-dom"
import { insert } from './service';



    export function AddDocument(){

        const navigate=useNavigate()
        const [errorName,setErrorName] = useState('')
        const[values,setValues]=useState({title:'',format:'',author:'',owner:'',validity:''})
  
  
        const set=(event)=>{
          setValues({...values,[event.target.name]:event.target.value})
        }

        const handleSubmit=async (e)=>{
          if (!values.title || !values.format || !values.author || !values.owner ||!values.validity){
            alert("please complete the details")
          }
          else{
            e.preventDefault()
            const resp = insert(values);
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

  return (
    <>
      <div className='leader'>
      <section className='container'>
        <header className='header'>Add Document</header>
      <Form className='form'
        labelCol={{
          span: 4,
        }}
        wrapperCol={{
          span: 14,
        }}
        layout="horizontal"
      >
        <div className='inputbox'>
        <Form.Item label="Title">
          <Input placeholder='Enter title' name='title' value={values.name} onChange={set} />
        </Form.Item>
        <Form.Item label="Format">
          <Input placeholder='Enter format' name='format' value={values.name} onChange={set} />
        </Form.Item>
        <Form.Item label="Author">
          <Input placeholder='Enter author name' name='author' value={values.name} onChange={set} />
        </Form.Item>
        <Form.Item label="Owner">
          <Input placeholder='Enter owner name' name='owner' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Validity">
          <Input placeholder='Enter Validity' type='date' name='validity' value={values.name} onChange={set}/>
        </Form.Item>
        <p style={{color:'red'}}>{errorName}</p>
        <Button type="primary" className='submit' onClick={handleSubmit}>Submit</Button>
        </div>
      </Form>
      </section>
      </div>
    </>
  );
};
