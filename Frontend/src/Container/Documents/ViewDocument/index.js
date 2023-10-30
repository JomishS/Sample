import './Document.scss';
import React, { useState,useEffect } from 'react';
import { Dropdown, Space,Button, Form, Popconfirm, Typography,Spin,InputNumber,Table,Input } from 'antd'
import { Link,useNavigate,useLocation } from "react-router-dom"
import { Tabledesign } from '../../../Component/Common/Table/Tabledesign';
import { fetchTheData,deleteid,update } from './service';
import {
  PlusOutlined,
  CloseOutlined,
  EditOutlined,
  DeleteOutlined,
} from "@ant-design/icons";




const onChange = (pagination, sorter, extra) => {
  console.log('params', pagination, sorter, extra);
};


export function Document() { 

  const navigate=useNavigate()
  let originData = [];

  const location=useLocation();
  var data
  data=location.state

  const [isCancelModalOpen, setIsCancelModalOpen] = useState(false);
  const [editingKey, setEditingKey] = useState('');

  useEffect(() => {
      fetch(data);
  },[isCancelModalOpen, editingKey] );

  const[docdetails,setDocDetails]=useState([])
  const fetch=(data)=>{
  const resp = fetchTheData(data);
  resp
  .then((res) => {   
    originData.push(res.data)
    console.log(originData)
    setDocDetails(originData)

  })
  .catch((err) => console.log(err));
}


  const [form] = Form.useForm();
  const isEditing = (record) => record.id === editingKey;
  const edit = (record) => {
    form.setFieldsValue({
      format: '',
      owner: '',
      validity: '',
      ...record,
    });
    setEditingKey(record.id);
  };


  const cancel = () => {
    setEditingKey('');
  };


  const handledelete=(id)=>{
    console.log(id)
    const resp = deleteid(id);
    resp
    .then((res) => {
      fetch()
    })
    .catch((err) => console.log(err));
  
  }
       

  
  const save = async (key) => {
      
    try {
     
      const row = await form.validateFields();
      const newData = [...docdetails];
      const index = newData.findIndex((item) => key === item.id);
      if (index > -1) {
        const item = newData[index];
        newData.splice(index, 1, {
          ...item,
          ...row,
        });
        const dataToUpdate=newData[index]
          const resp = update({key,dataToUpdate});
          resp
          .then((res) => {
            fetch()
          })
          .catch((err) => console.log(err));

        setDocDetails(newData);
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
      title: 'Id',
      dataIndex: 'id',
      align:'center',
      editable: false,
    },
    {
      title: 'Title',
      dataIndex: 'title',
      align:'center',
      editable: false,
    },
    {
      title: 'Format',
      dataIndex: 'format',
      align:'center',
      editable: true,
    },
    {
      title: 'Author',
      dataIndex: 'author',
      align:'center',
      editable: false,
    },
    {
      title: 'Owner',
      dataIndex: 'owner',
      align:'center',
      editable: true,
    },
    {
      title: 'Validity',
      dataIndex: 'validity',
      align:'center',
      editable: true,
    },

    {
      title: 'Update',
      dataIndex: 'update',
      align:'center',
      render: (_, record) => {
        const editable = isEditing(record);
        return editable ? (
          <span>
            <Popconfirm title="Are you sure" 
              onConfirm={() =>
              save(record.id)}
              onCancel={cancel}
              style={{ marginRight: 8 }}>
              <a>Save </a>
            </Popconfirm>

            <Popconfirm title="Sure to cancel?" onConfirm={cancel}>
              <a>Cancel</a>
            </Popconfirm>
          </span>
        ) : (
          <Typography.Link disabled={editingKey !== ''} onClick={() => navigate('/editdocuments',{state:record.id})}>
            <EditOutlined/>
          </Typography.Link>
        );
      },
    },
    {
      title: 'Delete',
      dataIndex: 'delete',
      align:'center',
      render: (_, record) => {
        return (
          <Popconfirm title="Sure to delete?"  onConfirm={() => handledelete(record.id)}>
          <a disabled={editingKey !== ''}>
            <DeleteOutlined/>
          </a>
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
            
            {!docdetails? (
            <Spin/>
            
          ):(
            <>
            <Form form={form} component={false}>
            <Tabledesign data={docdetails} columns={mergedColumns}/>
            </Form>       
 
          </>
          )}
       
        
      </div>

    </>
  );
}

