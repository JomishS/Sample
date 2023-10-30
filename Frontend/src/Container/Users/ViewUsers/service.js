import axios from "axios";


export const fetchTheData= async (pageNum,pageSize,filter,sortUrl)=>{
  pageNum--
const resp = await axios.get(`http://localhost:8085/users?limit=${pageSize}&page=${pageNum}&filter=${filter}&${sortUrl}`)
  return resp;
  }
  
    
export const deleteid = (props) => {
    const resp = axios.delete(`http://localhost:8085/users/${props}`)
  return resp;
};

export const fetchid= async (data)=>{
const resp = await axios.get(`http://localhost:8085/users/${data}`)
  return resp;
  }

export const update=(props)=>{
  console.log(props)
  console.log(props.id)
    const resp = axios.put(`http://localhost:8085/users/${props.id}`,props,{
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },})

    return resp
}




  
