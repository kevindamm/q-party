import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "About ?-Party",
  description: "a trivial side project",
  base: "/about/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '?-Party', link: '../' },
      { text: 'About', link: '/' }
    ],

    sidebar: [
      {
       // text: 'Examples',
       // items: [
       //   { text: 'Markdown Examples', link: '/markdown-examples' },
       //   { text: 'Runtime API Examples', link: '/api-examples' }
       // ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/kevindamm/q-party' }
    ]
  }
})
