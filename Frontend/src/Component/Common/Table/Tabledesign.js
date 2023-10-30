import { Table,Form } from "antd"
import { EditableCell } from "./Inputfield"
import React, { useContext, useState,useEffect} from 'react';
import {handleTableChange} from "../../../Container/Users/ViewUsers/util"
import { callFetchAllUsers } from "../../../Container/Users/ViewUsers";



export function Tabledesign(props){

  
  const data=props.data
  const columns=props.columns

  const [pageNum, setPageNum] = useState();
  const [pageSize, setPageSize] = useState();
  const [filters, setFilters] = useState(undefined);


  const cancel = () => {
    console.log("canceled")
  };
  

    return(
        <>
            <Table
            components={{
              body: {
                cell: EditableCell,
              },
            }}
            bordered
            dataSource={[...data]}
            columns={columns}
            rowKey={(record) => record._id}
            onChange={handleTableChange}
            pagination={{
              current: props.pageNum,
              pageSize: props.pageSize,
              total: props.total,
              onChange:props.onChange,
            }}
          /> 
          
            
        </>
    )
}

