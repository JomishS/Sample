import './Document.scss';
import React, { useState } from 'react';
import axios from 'axios';
import { Dropdown, Space,Button, Form, Popconfirm, Typography,Spin,InputNumber,Table,Input } from 'antd'
import { Link,useNavigate,useLocation } from "react-router-dom"




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




export function Document() {
  
  
  let originData = [];
  const location=useLocation();
  var data

  data=location.state
  console.log(data)

  const fetchTheData = ()=>{

    axios.get(`http://localhost:8085/documents/${data}`).then((res)=>{
      // console.log(res.data)
      originData.push(res.data)
      console.log(originData)
    setDocDetails(originData)
  },(err)=>{
    console.log(err)
      console.log("erorr in getting results inside fetch data")
  })
  }


  const navigate=useNavigate()
  const [form] = Form.useForm();
  const [editingKey, setEditingKey] = useState('');
  const[docdetails,setDocDetails]=useState(fetchTheData)
  

  const isEditing = (record) => record.doc_id === editingKey;
  const edit = (record) => {
    form.setFieldsValue({
      format: '',
      owner: '',
      validity: '',
      ...record,
    });
    setEditingKey(record.doc_id);
  };


  const cancel = () => {
    setEditingKey('');
  };

  const handledelete = (key) => {
    console.log(key)
    axios.delete(`http://localhost:8085/documents/${key}`,).then((res)=>{
      console.log(res.data)
      fetchTheData()
      setDocDetails(res.data)
  },(err)=>{
      console.log("erorr in getting results inside fetch data")
  })
  };



  const save = async (key) => {  
    try {
     
      const row = await form.validateFields();
      const newData = [...docdetails];
      const index = newData.findIndex((item) => key === item.doc_id);
      if (index > -1) {
        const item = newData[index];
        newData.splice(index, 1, {
          ...item,
          ...row,
        });
        const dataToUpdate=newData[index]
        setDocDetails(newData);

        axios.put(`http://localhost:8085/documents/${key}`,dataToUpdate,{
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
        setDocDetails(newData);
        setEditingKey('');
      }
    } catch (errInfo) {
      console.log('Validate Failed:', errInfo);
    }
  };
       


  const columns = [
    {
      title: 'doc_id',
      dataIndex: 'doc_id',
      editable: false,
      sorter: (a, b) => a.doc_id - b.doc_id,
      sortDirections: ['ascend', 'descend'],
      // defaultSortOrder: 'descend'
    },
    {
      title: 'title',
      dataIndex: 'title',
      editable: false,
      sorter: (a, b) => a.title.length - b.title.length,
      sortDirections: ['ascend', 'descend'],
      // defaultSortOrder: 'descend'
    },
    {
      title: 'format',
      dataIndex: 'format',
      editable: true,
    },
    {
      title: 'author',
      dataIndex: 'author',
      editable: false,
    },
    {
      title: 'owner',
      dataIndex: 'owner',
      editable: true,
      sorter: (a, b) => a.owner.length - b.owner.length,
      sortDirections: ['ascend', 'descend'],
    },
    {
      title: 'validity',
      dataIndex: 'validity',
      editable: true,
      sorter: (a, b) => a.validity.length - b.validity.length,
      sortDirections: ['ascend', 'descend'],
    },

    {
      title: 'Update',
      dataIndex: 'update',
      render: (_, record) => {
        const editable = isEditing(record);
        return editable ? (
          <span>
            <Popconfirm title="Are you sure" 
              onConfirm={() =>
              save(record.doc_id)}
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
        return (
          <Popconfirm title="Sure to delete?"  onConfirm={() => handledelete(record.doc_id)}>
          <a disabled={editingKey !== ''}>Delete</a>
        </Popconfirm>
        )
      }
    },
  //   {
  //     title: 'Documents',
  //     dataIndex: 'documents',
  //     render: (_, record) => {
  //       return(
  //       <div>
  //         <a disabled={editingKey !== ''} >
  //         View_documents
  //         </a>
  //       </div>
  //       );
  //     }
  //   },
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
      <h1 className='h1'>Document</h1>
      <div className='text'>
        <br /><br />
    
        <Link to ='/insertdoc'><Button type="primary">Add Documents</Button></Link><br/><br/>
        </div>

        <div className='table'>
            
             <Form form={form} component={false}>
          <Table
            components={{
              body: {
                cell: EditableCell,
              },
            }}
            bordered
            dataSource= {docdetails}
            // size="small"
            columns={mergedColumns}
            rowClassName="editable-row"
            pagination={{
              onChange: cancel,
            }}
          />
          </Form> 
        
      </div>

    </>
  );
}

