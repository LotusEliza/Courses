<template>
    <transition name="fade">
        <section>
        </section>
    </transition>
</template>

<script>
    export default {
        methods: {
            confirm(names) {
                this.$buefy.dialog.confirm({
                    cancelText: "Cancel",
                    message: 'Remove this items: '+names+'?',
                    onConfirm: () =>  this.$emit('clicked', names)
                })
            },
        }
    }
</script>

<style>
    /*************Transition******************/
    .fade-enter-active, .fade-leave-active {
        transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
        opacity: 0;
    }

</style>
