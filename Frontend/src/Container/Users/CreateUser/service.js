import axios from "axios";

export const insert=(props)=>{


    const resp = axios.post("http://localhost:8085/users",props,{
      headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
  },})
    return resp
}