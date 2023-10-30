import { PlusOutlined } from '@ant-design/icons';
import React, { useState } from 'react';
import './styles.scss'
import axios from 'axios';
import {Button,Form,Input} from 'antd';
import {useNavigate } from "react-router-dom"
import { insert } from './service';
const { TextArea } = Input;





  export function Createuser(){
      const navigate=useNavigate()
      const [errorName,setErrorName] = useState('')
      const[values,setValues]=useState({first_name:'',last_name:'',age:'',email:'',city:'',phone:'',birth_date:'',sex:'',country:'',doc_id:''})
      const set=(event)=>{
        setValues({...values,[event.target.name]:event.target.value})
      }


      const handleChange = (event) => {
        const { name, value } = event.target;
        const intValue = parseInt(value, 10); 
        setValues({ ...values, [name]: intValue });
      };
      

      const handleSubmit=async (e)=>{
        let index=values.doc_id
        console.log(index)

    if(!values.first_name || !values.last_name || !values.city || !values.phone || !values.sex || !values.country || !values.birth_date || !values.doc_id || !values.email || !values.age){
        alert("please complete the details")
      }
      else{
        e.preventDefault()
        console.log(values)


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
        <header className='header'>Add User</header>
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
        <Form.Item label="First Name">
          <Input placeholder='Enter first name' name='first_name' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Last Name">
          <Input placeholder='Enter last name' name='last_name' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Age">
          <Input placeholder='Enter age' type='number' name='age' value={values.age} onChange={handleChange}/>
        </Form.Item>
        <Form.Item label="Email">
          <Input type='email' placeholder='Enter email' name='email' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="City">
          <Input placeholder='Enter city' name='city' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Phone">
          <Input placeholder='Enter phone number' name='phone' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Birth Date">
          <Input placeholder='Enter birthdate' type='date' name='birth_date' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Gender">
          <Input placeholder='Enter city' name='sex' value={values.name} onChange={set}/>
        </Form.Item>
        
        <Form.Item label="Country">
          <Input placeholder='Enter country' name='country' value={values.name} onChange={set}/>
        </Form.Item>
        <Form.Item label="Document id">
          <Input placeholder='Enter document id' type='number' name='doc_id' value={values.doc_id} onChange={handleChange}/>
        </Form.Item>
        <br/>
        <p style={{color:'red'}}>{errorName}</p>
        <Button type="primary" className='submit' onClick={handleSubmit}>Submit</Button>
        </div>
      </Form>
      </section>
      </div>
    </>
  );
};

