<template>
  <v-container> 
    <v-card flat>
      <v-card-text>
        <v-row>
          <v-col>
            <v-text-field type="number" label="new process time in second" single-line dense hide-details v-model="second" />
          </v-col>
          <v-col>
            <v-btn color="primary" @click="addProc">
              Add Process
            </v-btn>
            <v-btn color="primary" @click="getItems">
              Refesh Monitor
            </v-btn>
          </v-col>
          <v-col>
            <v-checkbox label="Cont. check" v-model="continousCheck" />
          </v-col>
        </v-row>
        <div class="mt-2">
          Last check: {{ lastCheck }}
        </div>
      </v-card-text>
    </v-card>

    <div>
      <v-row>
        <v-col cols="3" sm="3" md="3" v-for="(item,idx) in items" :key="'item-'+idx">
          <v-card flat>
            <v-card-text>
              <div style="font-size:0.6em">{{ item._id }}</div>
              <h3>{{ item.source }}<br/>{{ item.ref }}</h3>
              <div>
                <v-chip x-small color="success" v-if="item.status=='Running'">Running</v-chip>
                <v-chip x-small color="primary" v-if="item.status=='Done'">Done</v-chip>
                <v-chip x-small color="error" v-if="item.status=='Error'">Error</v-chip>
              </div>
              <div style="font-size:0.8em">
                <b>Start:</b>
                <br/>
                {{ item.start }}
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>
export default {
  data () {
    return {
      lastCheck: new Date(),
      second: 0,
      items: [],
      continousCheck: true
    }
  },

  watch: {
    continousCheck (nv) {
      if (nv===true) this.keepRead()
    }
  },

  mounted () {
    //this.getItems()
    this.keepRead()
  },

  methods: {
    getItems (cb) {
      this.$axios.post('/process/gets',{}).then(
        r => {
          this.items = r.data
          if (cb) cb()
        },

        e => {
          //this.handleError(e)
          if (cb) cb()
        }
      )
    },

    addProc () {
      this.$axios.post('/process/add',{
        s: this.second
      }).then(
        r => {
          if (!this.continousCheck)this.getItems()
        },

        e => {
          this.handleError(e)
        }
      )
    },

    keepRead () {
      var self = this
      if (self.continousCheck===true) {
        this.getItems(setTimeout(function(){
          self.lastCheck = new Date()
          self.keepRead()
        }, 1000))
      }
    },

    handleError (e) {
      if (e.response && e.response.data) {
        alert(e.response.data)
        return
      }

      alert(e)
    }
  }
}
</script>
