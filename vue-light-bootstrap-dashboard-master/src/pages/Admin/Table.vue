<template>
  <section>
<!--    <b-field grouped group-multiline>-->
<!--      <button class="button field is-danger" @click="checkedRows = []"-->
<!--              :disabled="!checkedRows.length">-->
<!--        <b-icon icon="close"></b-icon>-->
<!--        <span>Clear checked</span>-->
<!--      </button>-->
<!--      <b-select v-model="checkboxPosition">-->
<!--        <option value="left">Checkbox at left</option>-->
<!--        <option value="right">Checkbox at right</option>-->
<!--      </b-select>-->
<!--    </b-field>-->

    <confirm ref="conf" @clicked="deleteFunc"></confirm>
    <b-tabs>
      <b-tab-item label="Table">
        <b-table
          :data="data"
          :columns="columns"
          :checked-rows.sync="checkedRows"
          :is-row-checkable="(row) => row.id !== 3 && row.id !== 4"
          checkable
          paginated
          per-page="10"
          :checkbox-position="checkboxPosition">

          <template slot="bottom-left">
            <b>Total checked</b>: {{ checkedRows.length }}
          </template>

        </b-table>
      </b-tab-item>
<!--      ---------------------------------------Buttons------------------------------>
      <b-field grouped group-multiline>
        <button class="button field is-info" @click="update">
          <b-icon icon="close"></b-icon>
          <span>Update</span>
        </button>
        <button class="button field is-info" @click="confirmRemove">
          <b-icon icon="close"></b-icon>
          <span>Delete</span>
        </button>
<!--        <button class="button field is-info" @click="add">-->
<!--          <b-icon icon="close"></b-icon>-->
<!--          <span>Add</span>-->
<!--        </button>-->
      </b-field>

      <b-tab-item label="Checked rows">
        <pre>{{ checkedRows }}</pre>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
import Confirm from "../../components/Confirm";

export default {
  name: "Table",
  components: {
    'confirm': Confirm,
  },
  props: {
    data: {
      type: Array
    },
    columns: {
      type: Array
    },
  },
  data() {
    const data = this.data;
    const columns = this.columns;
    return {
      checkboxPosition: 'right',
      checkedRows: [],
    }
  },
  methods:{
    update (event) {
      this.$emit('update', this.checkedRows )
    },
    deleteFunc (event) {
      this.$emit('delete', this.checkedRows )
    },
    add (event) {
      this.$emit('add')
    },
    async confirmRemove(){
      let names = [];
      try {
        for (let i = 0; i < this.checkedRows.length; i++) {
          if(this.checkedRows[i].Name){
            names.push(this.checkedRows[i].Name)
          }else {
            names.push(this.checkedRows[i].ID)
          }
        }
      } catch (error) {
        console.log('error')
      } finally {
        this.$refs.conf.confirm(names)
      }
    },
  }
}
</script>

<style>
.b-checkbox.checkbox {
  outline: none !important;
  display: inline-flex !important;
  align-items: center !important;
}
.pagination .pagination-next, .pagination .pagination-previous {
  margin-top: -12px !important;
}

</style>
