module.exports = {
    future: {
        purgeLayersByDefault: true,
    },
    corePlugins: {
        animation: false,
    },
    purge: [
        './index.html',
        './src/*.bs.js',
    ],
    theme: {
        extend: {

        },
    },
    variants: {},
    plugins: [],
}