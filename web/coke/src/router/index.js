import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/IndexView.vue'

import { useUserSession } from '@/stores/userSession'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'index',
      component: IndexView,
    },
    {
      path: '/nodes',
      name: 'nodes',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/NodesView.vue'),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/users',
      name: 'users',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/UsersView.vue'),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('../views/SettingsView.vue'),
      meta: {
        requiresAuth: true,
      },
      children: [
        {
          path: '',
          name: 'settings-general',
          component: () => import('../components/settings/SettingsGeneral.vue'),
        },
        {
          path: 'keys',
          name: 'settings-keys',
          component: () => import('../components/settings/SettingsKeys.vue'),
        },
        {
          path: 'node',
          name: 'settings-node',
          component: () => import('../components/settings/SettingsNode.vue'),
        },
        {
          path: 'dns',
          name: 'settings-dns',
          component: () => import('../components/settings/SettingsDns.vue'),
        },
      ],
    },
  ],
})

router.beforeEach((to, from, next) => {
  const userSession = useUserSession()
  if (to.meta.requiresAuth && !userSession.isLoggedIn) {
    next('/')
  } else {
    next()
  }
})

export default router
