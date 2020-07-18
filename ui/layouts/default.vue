<template>
  <v-app>
    <v-app-bar app height="30" flat dark>
        <b>Process Monitor</b>
    </v-app-bar>

    <v-main fluid style="background-color:#dfdfdf">
      <v-container>
        <nuxt/>
      </v-container>
    </v-main>
  </v-app>
</template>


<style>
body {
  font-family: Tahoma, Geneva, Verdana, sans-serif;
  font-size: 12px;
  font-weight: normal;
  background-color: #ddd;
}

h1,
h2,
h3,
h4 {
  font-weight: normal;
}

.breadcrumb {
  font-size: 1em;
  text-decoration: none;
  margin-bottom: 8px;
}

h1 {
  font-size: 1.6em;
  font-weight: normal;
}

.v-btn,
.v-tab {
  text-transform: none;
  letter-spacing: 0px;
}

.v-content {
  margin-top: 10px;
  margin-left: 10px;
  margin-right: 10px;
}

.v-list-item {
  min-height: 25px;
}

.v-list-item__title {
  font-size: 1em;
}

.v-card__subtitle, .v-card__text, .v-card__title{
  padding:10px;
  font-size:1em;
}

.v-card__title{
  padding-bottom:5px;
  font-size: 1.2em;
}

.v-card__text{
  padding-top:5px;
}

.col {
  padding-top: 2px;
  padding-bottom: 2px;
}

.v-text-field {
  padding-top: 8px;
  margin-top: 8px;
  font-size:1em;
}

.v-label{
  font-size:1em;
}

.v-input{
  font-size:1em;
}

.left {
  text-align: left;
}

.right {
  text-align: right;
}
</style>

