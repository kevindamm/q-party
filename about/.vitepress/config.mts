import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "About ?-Party",
  description: "a trivial side project",
  base: "/about/",

  srcExclude: [
    '**/README.md',
    'node_modules',
    'LICENSE.md'
  ],

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '?-Party', link: '/../' },
      { text: 'Schedule', link: '/schedule' },
      { text: 'Self-Host', link: '/self-host' },
      { text: 'About', link: '/' }
    ],

    sidebar: [
      {
        text: 'Gameplay',
        link: '/gameplay/',
        items: [
          { text: 'lobby', link: '/gameplay/lobby' },
          { text: 'match', link: '/gameplay/match' },
          { text: 'live', link: '/gameplay/live' },
          { text: 'async', link: '/gameplay/async' },
          { text: 'solo', link: '/gameplay/solo' }]},
      {
        text: 'Roles',
        link: '/roles/',
        items: [
          { text: 'host', link: '/roles/host' },
          { text: 'contestant', link: '/roles/contestant' },
          { text: 'spectator', link: '/roles/spectator' }]},
      {
        text: 'System',
        link: '/system/',
        items: [
          { text: 'chat', link: '/system/chat' },
          { text: 'buzzer', link: '/system/buzzer' },
          { text: 'speak', link: '/system/speak' },
          { text: 'judge', link: '/system/judge' },
          { text: 'quality', link: '/system/quality' }]},
      {
        text: 'Thanks',
        link: '/thanks/',
        items: [
          { text: 'invite', link: '/invite' },
          { text: 'support', link: '/support' },
          { text: 'contribute', link: '/contribute' }]}
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/kevindamm/q-party' }
    ]
  }
})
