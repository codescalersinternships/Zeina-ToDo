



<script lang="ts">

import { onMount } from "svelte";

import { writable, type Writable } from "svelte/store";


  type AddTodoType = (todo: string,id:string) => void
  type ToggleCompletedType = (event: MouseEvent) => void
  type TodosAmountType = number

 export let addTodo: AddTodoType
  export let toggleCompleted: ToggleCompletedType
  export let todosAmount: TodosAmountType

  let todo = ''
  let id=''
  // function handleSubmit() {
  //   addTodo(todo,id)
  //   todo = ''
  // }
  export const todos: Writable<any[]> = writable([]) 

 //base URL
 const baseURL = "http://localhost:4000/api/v1/list"

//Method to pull data
const getTodos = async () => {
    const response = await fetch("http://localhost:4000/api/v1/list")
    const data = await response.json()
    todos.set(data.list)
    
    // todo=await data
}

//Runs when component loads
onMount(()=>{
    getTodos()
})
let createTitle;
let createid;



  //create function for form submission
  
 
  const createTodo = async (event) => {
      event.preventDefault()
      await fetch(baseURL, {
          method: "post",
          headers: {
              "Content-Type":"application/json"
          },
          body: JSON.stringify({
              
              title: createTitle,
              // body: createBody
          })
      })
      getTodos()
      addTodo(createTitle,id)
      //refetch todos
     

      //reset form
  
      createTitle = ""
      todo = ''

  }
</script>

<form on:submit|preventDefault={createTodo}>
	{#if todosAmount > 0}
    <input
      on:click={toggleCompleted}
      type="checkbox"
      id="toggle-all"
      class="toggle-all"
    />
	  <label aria-label="Mark all as complete" for="toggle-all">
	    Mark all as complete
	  </label>
	{/if}

  <input
		bind:value={createTitle}
    id="new-todo"
    class="new-todo"
    placeholder="What needs to be done?"
    type="text"
    autofocus
  />
</form>



<style>
  .toggle-all {
    width: 1px;
    height: 1px;
    position: absolute;
    opacity: 0;
  }

  .toggle-all + label {
    position: absolute;
    font-size: 0;
  }

  .toggle-all + label:before {
    content: '❯';
    display: block;
    padding: var(--spacing-16);
    font-size: var(--font-24);
    color: var(--color-gray-58);
    transform: rotate(90deg);
  }

  .toggle-all:checked + label:before {
    color: var(--color-gray-28);
  }

  .new-todo {
    width: 100%;
    padding: var(--spacing-16);
    padding-left: 60px;
    font-size: var(--font-24);
    border: none;
    color: black;
    border-bottom: 1px solid var(--shadow-1);
  }
</style>