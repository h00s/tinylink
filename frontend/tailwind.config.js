/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontSize: {
        '4xl': '5rem',
      },
      colors: {
        'background-dark': '#011627',
        'background-light': '#fff',
        'dark-blue': '#011627',
        'light-green': '#99d19c',
        'dark-green': '#73ab84',
        'light-blue': '#ade1e5',
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        'light': {
          'base-content': '#000',
          primary: '#0fa5e9',
        },
      },
      {
        'dark': {
          'base-content': '#fff',
          primary: '#99d19c',
        },
      }
    ],
  },
}

// 011627