import './styles.scss';
import React, { useContext, useState,useEffect} from 'react';
import { Table,Button, Form, Popconfirm, Typography,Spin,InputNumber,Input, Pagination } from 'antd'
import { fetchTheData,update,deleteid } from './service';
import { Link,useNavigate} from "react-router-dom"
import { Tabledesign } from '../../../Component/Common/Table/Tabledesign';
import Filter from "cf-filter-component";
import { getFilterParams,isObjectEmpty,getFilterUrl,allUserConfig,getSortUrl} from './util';
import {
  EditOutlined,
  DeleteOutlined,
} from "@ant-design/icons";


const onChange = (pagination, sorter, extra) => {
  console.log('params', pagination, sorter, extra);
};
let sorters




const defaultStartPage = 1;
const defaultPageSize = 10;

export function Users() {

  let i=0
 let filterUrl,sortUrl,pageFilters
  const [pageNum, setPageNum] = useState(defaultStartPage);
  const [filters, setFilters] = useState(undefined);
  const [pageSize, setPageSize] = useState(defaultPageSize);
  const [loading, setLoading] = useState(true);
  const [totalCount, setTotalCount] = useState(0);


  const filtersCallback = (filterObj) => {
    setPageNum(1);
    pageFilters = getFilterParams(filterObj);
    if (filterObj && !isObjectEmpty(filterObj)) {
      setFilters(pageFilters);
    } else if (!isObjectEmpty(filters) && filters != "{}") {
      setFilters(undefined);
    }

    fetchAllUsers(pageNum, pageSize, pageFilters,sorters);
  };




 const [isCancelModalOpen, setIsCancelModalOpen] = useState(false);
  useEffect(() => {
    fetchAllUsers(defaultStartPage,defaultPageSize)
  },[isCancelModalOpen]);


const fetchAllUsers = (pageNum, pageSize, filters = {},sorter) => {
  setLoading(true);
  filterUrl = getFilterUrl(filters);
  const resp = fetchTheData(pageNum, pageSize, filterUrl,sorter);
  resp
    .then((res) => {
      setTotalCount(res.data.status.totalCount);
      convertIntoData(res.data.data)
    })
    .catch((err) => console.log(err));

  setLoading(false);


};




const convertIntoData=(data)=>{
    let new_data=
    data.map((elem,el)=>{ 
      return{
        id:elem.id,
        first_name: elem.first_name,
        last_name:elem.last_name,
        age:elem.age,
        email:elem.email,
        city:elem.city,
        phone:elem.phone,
        birth_date:elem.birth_date,
        sex:elem.sex,
        country:elem.country,
        doc_id:elem.doc_id
      };
    }) || [];
    setUserDetails(new_data)
  }
  
const[userdetails,setUserDetails]=useState([])
  
    const navigate=useNavigate()
    const [form] = Form.useForm();
    const [editingKey, setEditingKey] = useState('');
    const isEditing = (record) => record.id === editingKey;
  
  
    const edit = (record) => {
      form.setFieldsValue({
        age: '',
        email: '',
        city: '',
        phone: '',
        country: '',
        ...record,
      });
  
      setEditingKey(record.id);
    };
  
  
    const cancel = () => {
      setEditingKey('');
    };
  



    const handledelete=(id)=>{
      const resp = deleteid(id);
      resp
      .then((res) => {
        fetchAllUsers(pageNum, pageSize, pageFilters,sorters)
      })
      .catch((err) => console.log(err));
    
    }
  
  
    const save = async (key) => {
      
      try {
       
        const row = await form.validateFields();
        const newData = [...userdetails];
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
              // fetchAllUsers()
            })
            .catch((err) => console.log(err));

          setUserDetails(newData);
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
  
const handleSortChange=()=>{
  sortUrl=getSortUrl()
  fetchAllUsers(pageNum,pageSize,filterUrl,sortUrl)
}



  const columns = [
      
    {
      title: 'Id',
      dataIndex: 'id',
      align:'center',
      editable: false,
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    },
    {
      title: 'First Name',
      dataIndex: 'first_name',
      align:'center',
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    }, 
    
    {
      title: 'Last Name',
      dataIndex: 'last_name',
      align:'center',
      editable: false,
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    },
    {
      title: 'Age',
      dataIndex: 'age',
      align:'center',
      editable: true,
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    },
    {
      title: 'Email',
      dataIndex: 'email',
      align:'center',
      editable: true,
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    },
    {
      title: 'City',
      dataIndex: 'city',
      align:'center',
      editable: true,
    },
    {
      title: 'Phone',
      dataIndex: 'phone',
      align:'center',
      editable: true,
    },
    {
      title: 'Birth Date',
      dataIndex: 'birth_date',
      align:'center',
      editable: false,
      sorter: true, 
      sortDirections: ['ascend', 'descend'],
      onHeaderCell: (column) => ({
        onClick: () => handleSortChange(),
      }),
    },
    {
      title: 'Sex',
      dataIndex: 'sex',
      align:'center',
      editable: false,
    },
    {
      title: 'Country',
      dataIndex: 'country',
      align:'center',
      editable: true,
    },

    {
      title: 'Update',
      dataIndex: 'update',
      align:'center',
      render: (_, record) => {
        const editable = isEditing(record);
        // console.log(editable)
        return editable ? (
          <span>
            <Popconfirm title="Are you sure" onConfirm={() =>
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

          <Typography.Link disabled={editingKey !== ''} onClick={() => {
            console.log(record.id)
            navigate('./editusers',{state:record.id})
        }}>      
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
        return  (

          <Popconfirm title="Sure to delete?"  onConfirm={() => handledelete(record.id)}>
          <a disabled={editingKey !== ''}>
          <DeleteOutlined/>
          </a>
        </Popconfirm>
        );
      },
    },
    {
      title: 'Details',
      dataIndex: 'details',
      align:'center',
      render: (_, record) => {
        return(
        <div>
          <a disabled={editingKey !== ''} onClick={()=>{
            navigate('./documents',{state:record.doc_id})
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
         </div>
         <div className="all-vehicle-requests-container">
        <div>
         <Filter
          filters={allUserConfig}
          filtersCallback={filtersCallback}
        />
         </div>
          <div className="all-vehicle-requests-table">
          {!userdetails? (
            //  {loading ? (
            
            <Spin/>
            
          ):(
            <>
            <Form form={form} component={false}>
            <Tabledesign  data={userdetails} columns={mergedColumns} count={totalCount} pageNum={pageNum} pageSize={pageSize} total={totalCount} onChange={(pageNum, pageSize) => {
                setPageNum(pageNum);
                setPageSize(pageSize);
                fetchAllUsers(pageNum, pageSize, filters, sorters);
              }}/>
           
    
            </Form>       
 
          </>
          )}
        </div>
      </div>
     
    </>
  );
}


