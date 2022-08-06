//import axios from 'axios'



 export interface ITodo {
    id: string
    title: string
    done: boolean
  }

  // axios.get('http://localhost:4000/api/v1/list').then((res) =>{
  //     res.data.forEach((element:ITodo) => {
  //       const {id,text}=element
  //       console.log(id,text)
  //     });
  // })