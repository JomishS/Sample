import axios from "axios";


export const fetchTheData= (data)=>{
const resp = axios.get(`http://localhost:8085/documents/${data}`)
  return resp;
  }
  
    
export const deleteid = (props) => {
    const resp = axios.delete(`http://localhost:8085/documents/${props}`)
  return resp;
};


export const update=(props)=>{
    const resp = axios.put(`http://localhost:8085/documents/${props.id}`,props,{
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },})

    return resp
}




  
