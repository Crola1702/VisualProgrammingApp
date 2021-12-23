<template>
  <div id="App">
    <div class="container">
      <div class="row">
        <div id="numbers" class="col-6">
          <number-node @drag-node="dragNode"/>
        </div>
        <div id="arithmetic-operators" class="col-6 fondo-2">
        </div>
      </div>
      <div class="row">
        <button @click="peticion" class="btn btn-primary"></button>
      </div>
    </div>
  </div>
  <!-- <div id="App" class="container-fluid">
    <div class="row">
      <div class="col-md-12">
        <h1>Personas</h1>
      </div>
    </div>
    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="card-body">            
            <formulario-persona @add-persona="agregarPersona"/>
            <div v-if="!personas.lenght" class="alert alert-info" role="alert">
              No hay personas en la lista!
            </div>
            <tabla-personas :personas="personas" @delete-persona="eliminarPersona"/>
          </div>
        </div>
      </div>
    </div>
  </div> -->
</template>

<script>
// import TablaPersonas from '@/components/TablaPersonas.vue'
// import FormularioPersona from '@/components/FormularioPersona.vue'
import NumberNode from '@/components/NumberNode.vue'
import axios from 'axios';  

export default {
  name: 'App',
  components: {
    NumberNode
    // FormularioPersona,
    // TablaPersonas,
  },
  data() {
    return {
      personas: [
        {
          nombre: 'Jon',
          apellido: 'Doe',
          email: 'jondoe@gmail.com'
        },
        {
          nombre: 'Jane',
          apellido: 'Doe',
          email: 'janedoe@gmail.com'
        },
        {
          nombre: 'Mary',
          apellido: 'Doe',
          email: 'marydoe@gmail.com'
        },
        {
          nombre: 'Ken',
          apellido: 'Doe',
          email: 'kendoe@gmail.com'
        }
      ]
    }
  },
  methods: {
    agregarPersona(persona) {
      this.personas.push(persona);
    },
    eliminarPersona(email) {
      console.log(email);
      this.personas = this.personas.filter(persona => persona.email !== email);
    },

    dragNode(obj) {
      console.log(obj);
    },
    peticion() {
      axios.get('http://localhost:3000/nodes')
      .then(response => {
        console.log(response);
      })
      .catch(error => {
        console.log(error);
      });
    },
  },
}
</script>

<style>
#numbers {
  background-color: #630a0a;
}
#arithmetic-operators {
  background-color: #107922;
}
</style>
