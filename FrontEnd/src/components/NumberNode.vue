<template>
    <div id='number-node'>
        <div class="container-fluid">
            <div class="card border-dark text-white bg-primary mb-3" style="max-width: 13rem; max-height: 15rem;">
                <div class="row">
                    <div class="col-10">
                        <div class="card-body" >
                            <h5 class="card-title">Number</h5>
                            <div class="col">
                                <input type="text" v-model="number" class="form-control">
                            </div>
                        </div>
                    </div>
                    <div class="col-2" style="padding-top: 2rem">
                        <button v-bind:style="{ display: display }" type="button" class="edge-button" draggable="true" @dragstart="drag($event)"></button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'number-node',
        components: {
        },
        data() {
            return {
                id: '',
                number: 0,
                display: '',
            }
        },
        methods: {
            drag(event) {
                event.dataTransfer.dropEffect = 'move';
                event.dataTransfer.effectAllowed = 'true';
                event.dataTransfer.setData('nodeID', this.id);
                this.$emit('drag-node', {id: this.id, type: 'Number', event: event.dataTransfer});
            }
        }
    }
</script>

<style scoped>
    .edge-button {
        width: 30px;
        height: 30px;
        border-radius: 50%;
    }
</style>