<template>

    <div id="tabla-personas">
        <table class="table table-primary">
            <thead>
                <tr>
                    <th scope="col">Nombre</th>
                    <th scope="col">Apellido</th>
                    <th scope="col">Email</th>
                    <th scope="col">Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="persona in personas"  :key="persona.email">
                    <td>{{ persona.nombre }}</td>
                    <td>{{ persona.apellido }}</td>
                    <td>{{ persona.email }}</td>
                    <td>
                        <button class="btn btn-danger" @click="$emit('delete-persona', persona.email)">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>

</template>

<script>
export default {
    name: 'tabla-personas',
    props: {
        personas: Array,
    }
}
</script>

<style scoped>

</style>