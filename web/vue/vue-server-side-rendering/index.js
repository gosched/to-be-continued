const Vue = require('vue')
const server = require('express')()
const renderer = require('vue-server-renderer').createRenderer()

server.get('*', (req, res) => {
    const app = new Vue({
        data: {
            list: [{
                name: 'apple'
            }, {
                name: 'google'
            }]
        },
        template: `<ul><li v-for="item in list">{{item.name}}</li></ul>`
    })

    renderer.renderToString(app, (err, html) => {
        res.end(`
      <!DOCTYPE html>
      <html>
        <body>${html}</body>
      </html>
    `)
    })
})

console.log('http://127.0.0.1:8080/')
server.listen(8080)