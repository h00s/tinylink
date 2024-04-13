/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'background-dark': '#011627',
        'background-light': '#fff',
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
        },
      }
    ],
  },
}

// 011627