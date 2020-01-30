import Vue from "vue";
import VueRouter from "vue-router";
import Home from "@/Home.vue";
import agencyPageContainer from "@/components/agency/agencyPageContainer.vue";
import about from "@/components/about/about.vue";
import contact from "@/components/contact/contact.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/orgao/:agencyName",
    name: "agency",
    component: agencyPageContainer
  },
  {
    path: "/sobre",
    name: "sobre",
    component: about
  },
  {
    path: "/contato",
    name: "contato",
    component: contact
  }
];

const router = new VueRouter({
  routes
});

export default router;
