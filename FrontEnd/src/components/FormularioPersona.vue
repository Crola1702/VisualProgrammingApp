<template>
    <div id="formulario-persona">
        <form @submit.prevent="enviarFormulario">
            <div class="container-fluid">
                <div class="row">
                    <div class="col-md-4">
                        <div class="form-group">
                            <label>Nombre</label>
                            <input 
                                ref="nombre"
                                type="text" 
                                class="form-control" 
                                v-model="persona.nombre"
                                :class="{'is-invalid': procesando && nombreInvalido}"
                                @focus="resetEstado"
                                @keypress="resetEstado"
                            />
                        </div>
                    </div>
                    <div class="col-md-4">
                        <div class="form-group">
                            <label>Apellido</label>
                            <input 
                                type="text" 
                                class="form-control" 
                                v-model="persona.apellido"
                                :class="{'is-invalid': procesando && apellidoInvalido}"
                                @focus="resetEstado"
                            />
                        </div>
                    </div>
                    <div class="col-md-4">
                        <div class="form-group">
                            <label>Email</label>
                            <input 
                                type="email" 
                                class="form-control" 
                                v-model="persona.email"
                                :class="{'is-invalid': procesando && emailInvalido}"
                                @focus="resetEstado"
                            />
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-4">
                        <div class="form-group">
                            <button class="btn btn-primary">Guardar</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div v-if="error && procesando" class="alert alert-danger" role="alert">
                            Debes rellenar todos los campos!
                        </div>
                        <div v-if="correcto" class="alert alert-success" role="alert">
                            La persona ha sido agregada correctamente!
                        </div>
                    </div>
                </div>
            </div>
        </form>
    </div>
</template>

<script>
    export default {
        name: 'formulario-persona',
        data() {
            return {
                procesando: false,
                correcto: false,
                error: false,
                persona: {
                    nombre: '',
                    apellido: '',
                    email: ''
                }
            }
        },
        methods: {
            enviarFormulario() {
                this.procesando = true;
                this.resetEstado()


                if (this.nombreInvalido || this.apellidoInvalido || this.emailInvalido) {
                    this.error = true;
                    return;
                }

                this.$emit('add-persona', this.persona);
                this.$refs.nombre.focus();

                this.error = false;
                this.correcto = true;
                this.procesando = false;
                this.persona = {
                    nombre: '',
                    apellido: '',
                    email: ''
                }
            },
            resetEstado() {
                this.correcto = false;
                this.error = false;
            },
        },
        computed: {
            nombreInvalido() {
                return this.persona.nombre.length < 3;
            },
            apellidoInvalido() {
                return this.persona.apellido.length < 3;
            },
            emailInvalido() {
                return this.persona.email.length < 3;
            }
        }
    }
</script>

<style scoped>
    form {
        margin-bottom: 2rem;
    }
</style>
