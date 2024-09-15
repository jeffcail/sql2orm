<template>
  <div class="common-layout">
    <el-container>
      <!-- header -->
      <el-header class="header"> </el-header>

      <!-- main -->
      <el-main>
        <div class="main-header">
          <p class="title">SQL to ORM Converter</p>
          <p>
            This tool converts mysql create statements into Go type definitions.
          </p>
          <p>Other SQL expressions are ignored.</p>
          <p>
            Paste a create statement on the left and the equivalent Go type will
            be generated on the right,
          </p>
          <p>
            which you can paste into your program. The script has to make some
            assumptions,
          </p>
          <p>so check the output carefully!</p>
        </div>

        <div class="main-container-header">
          <el-radio-group v-model="radio">
            <el-radio :value="1">xorm</el-radio>
            <el-radio :value="2">gorm</el-radio>
          </el-radio-group>
        </div>
        <div class="main-container">
          <div class="main-container-left">
            <textarea
              class="textarea"
              v-model="sqlInput"
              placeholder="Enter SQL create table statement"
              @blur="convertSQL"
              ref="sqlTextarea"
              @keydown="handleKeydown"
            ></textarea>
          </div>

          <div class="main-container-right">
            <textarea class="textarea" id="output2">
                {{ outStruct }} 
            </textarea>
          </div>
        </div>
      </el-main>

      <!-- footer -->
      <el-footer class="footer"> 浙ICP备2021035698号-2 </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import axios from "axios";
import "../assets/css/home.css";
import { onMounted, ref } from "vue";

const sqlInput = ref("");
const outStruct = ref("");

const radio = ref(1);

onMounted(() => {
  // 1. 禁止使用右键菜单
  document.oncontextmenu = function (e) {
    e.returnValue = false;
  };

  // 2. 禁止鼠标选中
  document.onselectstart = function (e) {
    e.returnValue = false;
  };

  // 3. 禁止使用F12
  document.onkeydown = function (e: KeyboardEvent) {
    if (e.key == "F12") {
      e.preventDefault();
    }
  };
});

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === "a" || (e.ctrlKey && e.metaKey)) {
    e.preventDefault();
    const textarea = document.querySelector("textarea");
    if (textarea) {
      textarea.select();
    }
  }
};

// const handleKeydown2 = (e) => {
//   if (e.key === "a" || (e.ctrlKey && e.metaKey)) {
//     e.preventDefault();
//     const textarea = document.querySelector("#output2");
//     if (textarea) {
//       textarea.select();
//     }
//   }
// };

const convertSQL = async () => {
  // alert(radio.value);
  if (sqlInput.value == "") {
    outStruct.value = "Please fill in the SQL statement to be converted...";
    return;
  }

  const res = await axios.post("http://127.0.0.1:7892/gen", {
    sql: sqlInput.value,
    typ: radio.value,
  });
  outStruct.value = res.data.struct;
};
</script>

<style scoped>
</style>

