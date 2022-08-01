<!-- <script lang="ts">
  import Todos from '$root/components/Todos.svelte'
  import '$root/styles/global.css'
import PostList from './components/PostList.svelte'
</script>  -->

<Todos />

<!-- <input class="edit" type="text" autofocus />  -->

<script lang="ts">
  import {onMount} from 'svelte'
  import Todos from '$root/components/Todos.svelte'
  import '$root/styles/global.css'
  import {writable, type Writable} from "svelte/store";


export const todos: Writable<any[]> = writable([])


// export const fetchTodo = async () => {
//     const response =await fetch('http://localhost:4000/api/v1/list')
//     const data =await response.json()
//     //const {id, title, done}=data;
//    // document.getElementById('title').textContent= title;
//     //console.log(title)
//     todos.set(data.list)

//}
  //Variable to hold todos
  //  let todo = []

  // //base URL
  // const baseURL = "http://localhost:4000/api/v1/list"

  // //Method to pull data
  // const getTodos = async () => {
  //     const response = await fetch(baseURL)
  //     const data = await response.json()
  //     todos.set(data.list)
  // }

  // //Runs when component loads
  // onMount(()=>{
  //     getTodos()
  // })

  // //properties for create form
  // let createTitle;
  // let createBody;

  // //create function for form submission
  // const createTodo = async (event) => {
  //     event.preventDefault()
  //     await fetch(baseURL, {
  //         method: "post",
  //         headers: {
  //             "Content-Type":"application/json"
  //         },
  //         body: JSON.stringify({
  //             title: createTitle,
  //             // body: createBody
  //         })
  //     })

  //     //refetch todos
  //     getTodos()

  //     //reset form
  //     createTitle = ""
  //     // createBody = ""
  // }


  //     const deleteTodo = async (todo) => {
  //     event.preventDefault()
  //     await fetch(baseURL + "/" + todo.id, {
  //         method: "delete",
  //     })

  //     //refetch todos
  //     getTodos()
  // }



</script>




<main>

<!-- <h1>The Todo List</h1>
<hr>
<h2>Create a task</h2>
  <form on:submit={createTodo}>
      <input type="text" bind:value={createTitle}/>
       <input type="text" bind:value={createBody}/> -->
      <!-- <input type="submit" value="Create task"/>
  </form> -->
<!-- <hr>
<h2>Edit a Todo</h2>
  <form on:submit={updateTodo}>
      <input type="text" bind:value={editTitle}/>
      <input type="text" bind:value={editBody}/>
      <input type="submit" value="Update Todo"/>
  </form>
<hr> -->
<!-- <ul>
<h2>The Todos</h2>
{#each $todos as todo}
  <div>
      <h2>{todo.title}</h2>
       <h3>{todo.body}</h3> -->
      <!-- <button on:click={(e) => editSelect(todo)}>Edit</button> -->
      <!-- <button on:click={(e) => deleteTodo(todo)}>Delete</button>
  </div>
{/each}
</ul> --> 
</main>




<style>
 
  /* main {
    padding: 1em;
    max-width: 40em;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  } */
</style>
