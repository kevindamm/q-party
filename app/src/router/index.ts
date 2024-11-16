import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },

    // ASYNC ROUTES BELOW THIS LINE ///////////////////////////////////////////
    {
      path: '/g/:episode(\\d+)',
      component: () => import('../views/EpisodeView.vue'),
      props: true,
      children: [
        {
          path: '/g/:episode/:round(first|second)',
          component: () => import('../views/EpisodeRoundView.vue'),
          props: true,
          children: [{
            path: '/g/:episode/:round/host',
            component: () => import('../views/HostControlView.vue'),
            props: true,
            beforeEnter: (/* to, from */) => {
              // TODO if hosting and moving from a different route,
              //   ping the server for host permissions and the mod token,
              // TODO also confirm that current game should be closed.

              // return false if should cancel visiting `to` route
            },
          }],
        },
        {
          path: '/g/:episode/break',
          name: 'episode-break',
          component: () => import('../views/EpisodeBreakView.vue'),
          props: true,
        },
        {
          path: '/g/:episode/final-preview',
          name: 'final-preview',
          component: () => import('../views/FinalRoundPreview.vue'),
          props: true,
        },
        {
          path: '/g/:episode/final',
          component: () => import('../views/FinalRoundView.vue'),
          props: true,
          children: [{
            path: '/g/:episode/final/host',
            component: () => import('../views/HostFinalView.vue'),
            props: true,
            beforeEnter: (/* to, from */) => {
              // TODO if hosting and moving from a different route,
              //   ping the server for host permissions and the mod token,
              // TODO also confirm that current game should be closed.

              // return false if should cancel visiting `to` route
            },
          }],
        },
      ],
    }
  ],
})

export default router
