import axios from "axios";
import { fetchTheData } from "./service";

let sorters
export const allUserConfig = {
    id: {
      id: "id",
      label: "ID",
      placeholder: "Enter ID",
      type: "MULTI-TEXT-INPUT",
      remote: true,
    },
    first_name: {
      id: "first_name",
      label: "First name",
      placeholder: "Enter First Name",
      type: "MULTI-SELECT-REMOTE",
      remote: true,
      remoteConfig: firstRemoteConfig,
    },
    last_name: {
      id: "last_name",
      label: "Last name",
      placeholder: "Enter Last name",
      type: "MULTI-SELECT-REMOTE",
      remote: true,
      remoteConfig: lastRemoteConfig,
    },
    age: {
      id: "age",
      label: "Age",
      placeholder: "Enter Age",
      type: "MULTI-TEXT-INPUT",
      remote: true,
    },
    birth_date: {
    id: "birth_date",
    label: "Birth Date",
    type: "DATE-RANGE",
  },
    sex: {
      id: "sex",
      label: "Sex",
      placeholder: "Select Gender",
      type: "MULTI-SELECT",
      values: [
        {
          label: "Male",
          value: "male",
        },
        {
          label: "Female",
          value: "female",
        },
      ],
    },
  };



  export const getFilterParams = (filterList) => {
    let paramArr = [];
  
    Object.keys(filterList).map((key) => {
      if (key === "id") {
        let id = [];
        filterList[key]?.map((el) => {
          id.push(decodeURIComponent(el?.value));
        });
        paramArr.push({
          field: "id",
          operator: "eq",
          value: id,
          options: {
            ignoreCase: true,
          },
        });
      } else if (key === "first_name") {
        let srcType = [];
        filterList[key]?.map((el) => {
          srcType.push(decodeURIComponent(el?.value?.label));
        });
        paramArr.push({
          field: "first_name",
          operator: "eq",
          value: srcType,
          options: {
            ignoreCase: true,
          },
        });
      } else if (key === "last_name") {
        let srcType = [];
        filterList[key]?.map((el) => {
          srcType.push(decodeURIComponent(el?.value?.label));
        });
        paramArr.push({
          field: "last_name",
          operator: "eq",
          value: srcType,
          options: {
            ignoreCase: true,
          },
        });
      } else if (key === "sex") {
        let mileType = [];
        filterList[key]?.map((el) => {
          mileType.push(decodeURIComponent(el?.value?.value));
        });
        paramArr.push({
          field: "sex",
          operator: "eq",
          value: mileType,
          options: {
            ignoreCase: true,
          },
        });
      } else if (key === "age") {
        let id = [];
        filterList[key]?.map((el) => {
          id.push(decodeURIComponent(el?.value));
        });
        paramArr.push({
          field: "age",
          operator: "eq",
          value: id,
          options: {
            ignoreCase: true,
          },
        });
      } else if (key === "birth_date") {
        let values = [];
        filterList[key].map((filter) => {
          values.push(filter.value.from, filter.value.to);
        });
       
        paramArr.push({
          field: "birth_date",
          operator: "between",
          value: values,
          options: {
            ignoreCase: true,
          },
        });
      }
    });
    return paramArr;
  };



  export const isObjectEmpty = (obj) => {
    for (var i in obj) return false;
    return true;
  };


  export const getFilterUrl = (filters) => {
    let url = "";
    if (filters?.length >= 1) {
      filters?.map((el) => {
        if (el?.field == "id") {
          if (url?.length >= 1) {
            url = url + `&id=${decodeURIComponent(el?.value.join(","))}`;
          } else {
            url=`id.in:${decodeURIComponent(el?.value.join(","))}`
          }
        } else if (el?.field == "first_name") {
          if (url?.length >= 1) {
            url =
              url + `&first_name=${decodeURIComponent(el?.value.join("|"))}`;
          } else {
            url=`first_name.in:${decodeURIComponent(el?.value.join(","))}`
            console.log(url)
          }
        } else if (el?.field == "last_name") {
          if (url?.length >= 1) {
            url =
              url +
              `&last_name=${decodeURIComponent(el?.value.join("|"))}`;
          } else {
            url=`last_name.in:${decodeURIComponent(el?.value.join(","))}`
          }
        } else if (el?.field == "age") {
            if (url?.length >= 1) {
              url = url + `&age=${decodeURIComponent(el?.value.join(","))}`;
            } else {
            url=`age.in:${decodeURIComponent(el?.value.join(","))}`
            }
          } else if (el?.field == "sex") {
          if (url?.length >= 1) {
            url = url + `&sex=${decodeURIComponent(el?.value.join("|"))}`;
          } else {
            url=`sex.eq:${decodeURIComponent(el?.value.join(","))}`
          }
        } else  if (el?.field == "birth_date") {
          if (url?.length >= 1) {
            url = url +`&birth_date=${decodeURIComponent(el?.value.join("|"))}`;
            
          } else {
            url=`birth_date.gte:${el.value[0]},birth_date.lte:${el.value[1]}`
          }
        }
      });
    }
    // console.log(url)
    return decodeURIComponent(url);
  };



 

  


  export async function firstRemoteConfig(searchInput) {
    let path = "http://localhost:8085/users?filter=";
    try {
      if (typeof searchInput == "string") {
        path = `${path}first_name.eq:${searchInput}`;
      }
      const response = await axios.get(path);
      const data = response.data;
      const distinctFirstNames = new Set();
  
      let arr = [];
      if (data?.status.totalCount) {

        data.data.forEach((item) => {
          const firstName = item.first_name;
          if (!distinctFirstNames.has(firstName)) {
            let obj = {};
            obj.label = firstName;
            obj.value = item.id;
            arr.push(obj);
            distinctFirstNames.add(firstName);
          }
        });
      }
      return arr;
    } catch (error) {
      throw error;
    }
  }



  // export async function firstRemoteConfig(searchInput) {
  //   let path = "http://localhost:8085/users?filter=";
  //   try {
  //     if (typeof searchInput == "string") {
  //       path = `${path}first_name.eq:${searchInput}`;
  //       console.log(path)
  //     }
  //     const response = await axios.get(path);
  //     const data = response.data;
  //     console.log(data.status.totalCount)
  //     const distinctFirstNames = new Set();
  
  //     let arr = [];
  //     // console.log(data.length)
  //     if (data?.length) {

  //       data.forEach((item) => {
  //         const firstName = item.first_name;
  //         if (!distinctFirstNames.has(firstName)) {
  //           let obj = {};
  //           obj.label = firstName;
  //           obj.value = item.id;
  //           arr.push(obj);
  //           distinctFirstNames.add(firstName);
  //         }
  //       });
  //     }
  //     return arr;
  //   } catch (error) {
  //     throw error;
  //   }
  // }

  
  export async function lastRemoteConfig(searchInput) {
    let path = "http://localhost:8085/users?filter=";
    try {
      if (typeof searchInput == "string") {
        path = `${path}last_name.eq:${searchInput}`;
      }
      const response = await axios.get(path);
      const data = response.data;

      const distinctLastNames = new Set();
  
      let arr = [];
      if (data?.status.totalCount) {
        data.data.forEach((item) => {
          const lastName = item.last_name;
          if (!distinctLastNames.has(lastName)) {
            let obj = {};
            obj.label = lastName;
            obj.value = item.id;
            arr.push(obj);
            distinctLastNames.add(lastName);
          }
        });
      }
      return arr;
    } catch (error) {
      throw error;
    }
  }





export const handleTableChange = (pagination,filters, sorter) => {
    sorters=sorter
  };



export const getSortUrl=()=>{
  let url=""
  if(sorters?.order){
  if (sorters.order=='ascend'){
    url=`sortBy=${sorters.column.dataIndex}.asc`
    // console.log(url)
  } else if(sorters.order=='descend'){
    url=`sortBy=${sorters.column.dataIndex}.desc`
  } else{
    url=""
  }
}
  return decodeURIComponent(url);


}