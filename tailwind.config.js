/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      cloak: { raw: '(display: none !important;)' },
    },
  },
  plugins: [],
}

