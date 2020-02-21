const Component = Vue.component('temp', {
    data() {
        return {
            activeName: 'model',
            tempContent: null
        };
    },
    mounted() {
        this.handleReadTemp(this.activeName)
    },
    watch: {
        activeName(n, o) {
            if (n === o) {
                return
            }
            this.handleReadTemp(n)
        }
    },
    methods: {
        handleReadTemp(tempName) {
            axios.get('/api/v1/temp?temp_name=' + tempName).then(resp => {
                var data = resp.data.data
                console.log(data)
                if (resp.data.code !== 0) {
                    this.$message.error(data)
                } else {
                    this.tempContent = data
                }
            })
        }
    },
    template: `
        <div>
            <el-card class="box-card">
              <p>说明：</p>
              <p style="text-indent: 2em">structName      ==>   大驼峰结构体名称</p>
              <p style="text-indent: 2em">tableName       ==>   表名</p>
              <p style="text-indent: 2em">tableInfos      ==>   表信息</p>
              <p style="text-indent: 4em">CamelName       ==>   大驼峰字段名称</p>
              <p style="text-indent: 4em">GoType          ==>   Go类型</p>
              <p style="text-indent: 4em">ColumnName      ==>   数据库字段名</p>
              <p style="text-indent: 4em">ColumnComment   ==>   字段注释</p>
                
            </el-card>
            <el-tabs v-model="activeName">
                <el-tab-pane label="model" name="model">
                    <el-input type="textarea" autosize v-model="tempContent"></el-input>
                    <el-button>保存</el-button>
                </el-tab-pane>
            </el-tabs>
        </div>
    `
})

const Create = Vue.component('create', {
    data() {
        return {
            label_width: "100px",
            dbForm: {
                driver: "mysql",
                host: "127.0.0.1",
                port: "3306",
                db_name: "",
                username: "root",
                password: "root",
                extras: [],
                table_names: [],
            },
            options: [],
            dbFormRules: {
                driver: [
                    {
                        required: true,
                        message: "驱动必填",
                        trigger: "blur"
                    }
                ],
                host: [
                    {
                        required: true,
                        message: "主机必填",
                        trigger: "blur"
                    }
                ],
                port: [
                    {
                        required: true,
                        message: "端口号必填",
                        trigger: "blur"
                    }
                ],
                db_name: [
                    {
                        required: true,
                        message: "数据库名称必填",
                        trigger: "blur"
                    }
                ],
                username: [
                    {
                        required: true,
                        message: "用户名必填",
                        trigger: "blur"
                    }
                ],
                password: [
                    {
                        required: true,
                        message: "密码必填",
                        trigger: "blur"
                    }
                ]
            }
        };
    },
    mounted() {
        axios.post('/api/v1/db/tables', this.dbForm).then(resp => {
            this.options = resp.data.data;
        });
    },
    computed: {
        newDBName() {
            return this.dbForm.db_name
        }
    },
    watch: {
        newDBName(new_db_name, old) {
            setTimeout(() => {
                axios.post('/api/v1/db/tables', this.dbForm).then(resp => {
                    this.options = resp.data.data;
                });
            }, 2000)
        }
    },
    methods: {
        resetForm(formName) {
            this.$refs[formName].resetFields();
            this.dbForm.extras = [];
        },
        removeExtra(item) {
            let index = this.dbForm.extras.indexOf(item);
            if (index !== -1) {
                this.dbForm.extras.splice(index, 1);
            }
        },
        addExtra() {
            this.dbForm.extras.push({
                index: Date.now(),
                key: "",
                value: ""
            });
        },
        generate(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    window.location.href = '/api/v1/create?' +
                        'driver=' + this.dbForm.driver +
                        '&username=' + this.dbForm.username +
                        '&password=' + this.dbForm.password +
                        '&host=' + this.dbForm.host +
                        '&port=' + this.dbForm.port +
                        '&db_name=' + this.dbForm.db_name +
                        '&extras=' + this.dbForm.extras +
                        '&table_names=' + this.dbForm.table_names.join(",")
                } else {
                    return false;
                }
            })
        },
        handleClick(tab, event) {
            console.log(tab, event);
        }
    },
    template: `<el-form ref="dbForm" :model="dbForm" :rules="dbFormRules">
            <el-form-item label="驱动" :label-width="label_width" prop="driver">
                <el-input v-model="dbForm.driver" type="text" placeholder="请输入驱动名称" clearable disabled/>
            </el-form-item>
            <el-form-item label="主机" :label-width="label_width" prop="host">
                <el-input v-model="dbForm.host" placeholder="请输入主机" clearable/>
            </el-form-item>
            <el-form-item label="端口号" :label-width="label_width" prop="port">
                <el-input v-model="dbForm.port" placeholder="请输入端口号" clearable/>
            </el-form-item>
            <el-form-item label="数据库名称" :label-width="label_width" prop="db_name">
                <el-input v-model="dbForm.db_name" placeholder="请输入数据库名称" clearable/>
            </el-form-item>
            <el-form-item label="用户名" :label-width="label_width" prop="username">
                <el-input v-model="dbForm.username" placeholder="请输入用户名" clearable/>
            </el-form-item>
            <el-form-item label="密码" :label-width="label_width" prop="password">
                <el-input v-model="dbForm.password" placeholder="请输入密码" show-password clearable/>
            </el-form-item>

            <el-form-item label="额外参数" :label-width="label_width">
                <el-form-item v-for="ext in dbForm.extras" :key="ext.index"
                              :rules="{required: true, message: '键值不能为空', trigger: 'blur'}">
                    <el-input v-model="ext.key" placeholder="键" clearable style="width: 30%"/>
                    <el-input v-model="ext.value" placeholder="值" clearable style="width: 30%"/>
                    <el-button @click.prevent="removeExtra(ext)">删除</el-button>
                </el-form-item>
                <el-button @click="addExtra()">新增</el-button>
            </el-form-item>
            <el-form-item label="表" :label-width="label_width">
                <el-select v-model="dbForm.table_names" multiple filterable placeholder="请选择表,(多表默认所有表)"
                           style="width:250px">
                    <el-option
                            v-for="item in options"
                            :key="item.name"
                            :label="item.comment"
                            :value="item.name"
                    />
                </el-select>
            </el-form-item>
            <el-form-item :label-width="label_width">
                <el-button @click="resetForm('dbForm')">重置</el-button>
                <el-button @click="generate('dbForm')">生成</el-button>
            </el-form-item>
        </el-form>`
})

