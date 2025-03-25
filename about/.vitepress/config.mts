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
      { text: 'Questions', link: '/questions' },
      { text: 'Categories', link: '/category' },
      { text: 'Calendar', link: '/catwhen' },
      { text: 'Matches', link: '/match' },
      { text: 'About', link: '/' }
    ],

    sidebar: [
      {
        text: 'Gameplay',
        items: [
          { text: 'lobby', link: '/lobby' },
          { text: 'match', link: '/match' },
          { text: 'live', link: '/live' },
          { text: 'async', link: '/async' },
          { text: 'solo', link: '/solo' }]},
      {
        text: 'Roles',
        items: [
          { text: 'host', link: '/host' },
          { text: 'contestant', link: '/contestant' },
          { text: 'spectator', link: '/spectator' }]},
      {
        text: 'Tech',
        items: [
          { text: 'chat', link: '/chat' },
          { text: 'buzzer', link: '/buzzer' },
          { text: 'speak', link: '/speak' },
          { text: 'check', link: '/check' }]},
      {
        text: 'Thanks',
        items: [
          { text: 'invite', link: '/invite' },
          { text: 'support', link: '/support' },
          { text: 'contribute', link: '/contribute' },
          { text: 'contact', link: '/contact' }]}
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/kevindamm/q-party' }
    ]
  }
})
