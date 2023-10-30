import axios from "axios";

export const insert=(props)=>{
    console.log(props)
    // console.log(props.dataToUpdate)
    const resp = axios.post("http://localhost:8085/documents",props,{
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },})

    return resp
}