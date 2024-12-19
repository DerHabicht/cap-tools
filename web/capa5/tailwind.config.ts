import type { Config } from 'tailwindcss'
import colors from 'tailwindcss/colors'

export default <Partial<Config>>{
    plugins: [require('tailwindcss-primeui')],
    theme: {
        colors: {
            symbolBlue: '#001871',
            silverGray: '#9EA2A2',
            scarletRed: '#BA0C2F',
            afYellow: '#FFCD00',
            black: '#000000',
            white: '#FFFFFF',
            gray: '#666666',
        },
    },
    darkMode: 'selector',
}
