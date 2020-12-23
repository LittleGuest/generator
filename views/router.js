const router = new VueRouter({
    routes: [{
            path: '',
            component: Create
        },
        {
            path: '/temp',
            component: Temp
        },
        {
            path: '/create',
            component: Create
        }
    ]
});