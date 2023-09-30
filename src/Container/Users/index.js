import './Users.scss';
import React, { useState} from 'react';
import axios from 'axios';
import { Dropdown, Table,Button, Form, Popconfirm, Typography,Spin,InputNumber,Input } from 'antd'
import { DownOutlined } from '@ant-design/icons';
import { Link,useNavigate} from "react-router-dom"






const onChange = (pagination, sorter, extra) => {
  console.log('params', pagination, sorter, extra);
};

const EditableCell = ({
  editing,
  dataIndex,
  title,
  inputType,
  record,
  index,
  children,
  ...restProps
}) => {
  const inputNode = inputType === 'number' ? <InputNumber /> : <Input />;
  return (
    <td {...restProps}>
      {editing ? (
        
        <Form.Item
          name={dataIndex}
          style={{
            margin: 0,
          }}
          rules={[  
            {
              required: true,
              message: `Please Input ${title}!`,
            },
          ]}
        >
          {inputNode}
        </Form.Item>
      ) : (
        children
      )}
    </td>
  );
};




export function Users() {


const fetchTheData=()=>{
  axios.get('http://localhost:8085/users').then((res)=>{
    convertIntoData(res.data)
},(err)=>{
    console.log("erorr in getting results inside fetch data")
})
}


const convertIntoData=(data)=>{
  let new_data=
  data.map((elem,el)=>{ 
    return{
      user_id:elem.user_id,
      first_name: elem.first_name,
      last_name:elem.last_name,
      age:elem.age,
      email:elem.email,
      city:elem.city,
      phone:elem.phone,
      birth_date:elem.birth_date,
      sex:elem.sex,
      country:elem.country,
      document_id:elem.document_id
    };
  }) || [];
  setUserDetails(new_data)
}



  const navigate=useNavigate()
  const [form] = Form.useForm();
  const [editingKey, setEditingKey] = useState('');
  const[userdetails,setUserDetails]=useState(fetchTheData)
  const isEditing = (record) => record.user_id === editingKey;


  const edit = (record) => {
    form.setFieldsValue({
      age: '',
      email: '',
      city: '',
      phone: '',
      country: '',
      ...record,
    });

    setEditingKey(record.user_id);
  };


  const cancel = () => {
    setEditingKey('');
  };

  const handledelete = (key) => {
    axios.delete(`http://localhost:8085/users/${key}`,).then((res)=>{
      fetchTheData()
  },(err)=>{
      console.log("erorr in getting results inside fetch data")
  })
  };


  const save = async (key) => {
    
    try {
     
      const row = await form.validateFields();
      const newData = [...userdetails];
      const index = newData.findIndex((item) => key === item.user_id);
      if (index > -1) {
        const item = newData[index];
        newData.splice(index, 1, {
          ...item,
          ...row,
        });
        const dataToUpdate=newData[index]
        setUserDetails(newData);


        axios.put(`http://localhost:8085/users/${key}`,dataToUpdate,{
          headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
      },}).then((res)=>{
      },(err)=>{
          console.log("erorr in getting results inside fetch data")
      })

        setEditingKey('');
      } else {
        newData.push(row);
        setUserDetails(newData);
        setEditingKey('');
      }
    } catch (errInfo) {
      console.log('Validate Failed:', errInfo);
    }
  };



  const columns = [
     {
      title: 'user_id',
      dataIndex: 'user_id',
      editable: false,
      sorter: (a, b) => a.user_id - b.user_id,
      sortDirections: ['ascend', 'descend'],
      // defaultSortOrder: 'ascend'
    },
    {
      title: 'first_name',
      dataIndex: 'first_name',
      editable: false,
      sorter: (a, b) => a.first_name.length - b.first_name.length,
      sortDirections: ['ascend', 'descend'],
      // defaultSortOrder: 'ascend'
    },
    {
      title: 'last_name',
      dataIndex: 'last_name',
      editable: false,
      sorter: (a, b) => a.last_name.length - b.last_name.length,
      sortDirections: ['ascend', 'descend'],
    },
    {
      title: 'age',
      dataIndex: 'age',
      editable: true,
      sorter: (a, b) => a.age - b.age,
      sortDirections: ['ascend', 'descend'],
    },
    {
      title: 'email',
      dataIndex: 'email',
      editable: true,
      sorter: (a, b) => a.email.length - b.email.length,
      sortDirections: ['ascend', 'descend'],
    },
    {
      title: 'city',
      dataIndex: 'city',
      editable: true,
    },
    {
      title: 'phone',
      dataIndex: 'phone',
      editable: true,
    },
    {
      title: 'birth_date',
      dataIndex: 'birth_date',
      editable: false,
      sorter: (a, b) => a.birth_date.length - b.birth_date.length,
      sortDirections: ['ascend', 'descend'],
    },
    {
      title: 'sex',
      dataIndex: 'sex',
      editable: false,
    },
    {
      title: 'country',
      dataIndex: 'country',
      editable: true,
    },

    {
      title: 'Update',
      dataIndex: 'update',
      render: (_, record) => {
        const editable = isEditing(record);

        return editable ? (
          <span>
            <Popconfirm title="Are you sure" onConfirm={() =>
              save(record.user_id)}
              onCancel={cancel}
              style={{ marginRight: 8 }}>
              <a>Save </a>
            </Popconfirm>

            <Popconfirm title="Sure to cancel?" onConfirm={cancel}>
              <a>Cancel</a>
            </Popconfirm>
          </span>
        ) : (

          <Typography.Link disabled={editingKey !== ''} onClick={() => edit(record)}>
            Edit
          </Typography.Link>
        );
      },
    },
    {
      title: 'Delete',
      dataIndex: 'delete',
      render: (_, record) => {
        return  (

          <Popconfirm title="Sure to delete?"  onConfirm={() => handledelete(record.user_id)}>
          <a disabled={editingKey !== ''}>Delete</a>
        </Popconfirm>
        );
      },
    },
    {
      title: 'Documents Details',
      dataIndex: 'documents details',
      render: (_, record) => {
        return(
        <div>
          <a disabled={editingKey !== ''} onClick={()=>{
            navigate('./documents',{state:record.document_id})
          }}>
          View
          </a>
        </div>
        );
      }
    },

  ];


  const mergedColumns = columns.map((col) => {
    if (!col.editable) {
      return col;
    }
    return {
      ...col,
      onCell: (record) => ({
        record,
        inputType: col.dataIndex === 'age' ? 'number' : 'text',
        dataIndex: col.dataIndex,
        title: col.title,
        editing: isEditing(record),
      }),
    };
  });

  return (
    
    <>

      <h1 className='h1'>USERS</h1>
      <div className='text'>
        <br /><br />
               


         <Link to='/createuser'><Button type="primary">CreateUser</Button></Link>
         </div><br/>
       

        
      <div className='table'>
       
       {/* <Tabledesign data={userdetails} columns={columns} form={form} editingKey={editingKey}/> */}

          {!userdetails? (
            <Spin/>
            
          ):(
            <>
            
            <Form form={form} component={false}>
          <Table
            components={{
              body: {
                cell: EditableCell,
              },
            }}
            bordered
            dataSource={[...userdetails]}
            columns={mergedColumns}
            rowClassName="editable-row"
            pagination={{
              onChange: cancel,
            }}
          />
          </Form>
        
 
          </>
          )}
        
      </div>
     
    </>
  );
}

