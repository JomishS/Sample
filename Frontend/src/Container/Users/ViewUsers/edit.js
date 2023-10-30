import React, { useState, useEffect, useCallback } from 'react'
import { Button,Form,Input,Row,Col,Divider } from 'antd';
import { update } from './service';
import {useNavigate,useLocation } from "react-router-dom"
import { fetchid } from './service';


export function EditUsers(){

    const location=useLocation();
    const [isCancelModalOpen, setIsCancelModalOpen] = useState(false);
    const[values,setValues]=useState({first_name:'',last_name:'',age:'',email:'',city:'',phone:'',birth_date:'',sex:'',country:'',doc_id:''})
    var data
    data=location.state


    useEffect(() => {
        fetch(data);
    },[isCancelModalOpen] );


    const fetch = (data) => {
    const resp = fetchid(data);
    resp
      .then((res) => {
        console.log(res.data)
        const updatedValues = { ...values }; 
        updatedValues.id=res.data.id
        updatedValues.first_name = res.data.first_name;
        updatedValues.last_name = res.data.last_name;
        updatedValues.age = res.data.age;
        updatedValues.email = res.data.email;
        updatedValues.city = res.data.city;
        updatedValues.phone = res.data.phone;
        updatedValues.birth_date = res.data.birth_date;
        updatedValues.sex = res.data.sex;
        updatedValues.country = res.data.country;
        updatedValues.doc_id = res.data.doc_id;
  
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


      const handleChange = (event) => {
        const { name, value } = event.target;
        const intValue = parseInt(value, 10); 
        setValues({ ...values, [name]: intValue });
      };
      

      const handleSubmit=async (e)=>{
        let index=values.id
        console.log(data)

    if(!values.first_name || !values.last_name || !values.city || !values.phone || !values.sex || !values.country || !values.birth_date || !values.doc_id || !values.email || !values.age){
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
        <header className='header'>Edit User</header>
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
          <Input name='first_name' value={values.first_name} onChange={set} disabled={true}/>
        </Form.Item>
        <Form.Item label="Last Name">
          <Input name='last_name' value={values.last_name} onChange={set} disabled={true}/>
        </Form.Item>
        <Form.Item label="Age">
          <Input  type='number' name='age' value={values.age} onChange={handleChange}/>
        </Form.Item>
        <Form.Item label="Email">
          <Input type='email'  name='email' value={values.email} onChange={set}/>
        </Form.Item>
        <Form.Item label="City">
          <Input  name='city' value={values.city} onChange={set}/>
        </Form.Item>
        <Form.Item label="Phone">
          <Input  name='phone' value={values.phone} onChange={set}/>
        </Form.Item>
        <Form.Item label="Birth Date">
          <Input  type='date and time' name='birth_date' value={values.birth_date} onChange={set} disabled={true}/>
        </Form.Item>
        <Form.Item label="Gender">
          <Input  name='sex' value={values.sex} onChange={set} disabled={true}/>
        </Form.Item>
        
        <Form.Item label="Country">
          <Input  name='country' value={values.country} onChange={set}/>
        </Form.Item>
        <Form.Item label="Document id">
          <Input  type='number' name='doc_id' value={values.doc_id} onChange={handleChange} disabled={true}/>
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