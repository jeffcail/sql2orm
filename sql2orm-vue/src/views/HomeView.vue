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
            <!-- 模板部分 -->
            <div class="output-wrapper">
              <pre class="code-output" v-text="displayedOutput"></pre>

              <el-button
                v-if="outStruct && !outStruct.includes('错误')"
                size="small"
                type="primary"
                class="copy-btn"
                @click="copyCode"
                >复制代码</el-button
              >
            </div>
          </div>
        </div>
      </el-main>

      <!-- footer -->
      <el-footer class="footer"> ©2020 mazezen mazezen24@gmail.com </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import axios from "axios";
import "../assets/css/home.css";
import { onMounted, ref } from "vue";
import { ElMessage } from "element-plus";
import { watch } from "vue";
import { computed } from "vue";

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

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(outStruct.value);
    ElMessage.success("已复制到剪贴板！");
  } catch (err) {
    ElMessage.error("复制失败");
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

  const res = await axios.post("/gen", {
    sql: sqlInput.value,
    typ: radio.value,
  });
  outStruct.value = res.data.struct;
};

watch(
  radio,
  (newVal, oldVal) => {
    if (newVal !== oldVal) {
      sqlInput.value = ""; // 清空左侧 SQL 输入框
      outStruct.value = ""; // 清空右侧生成的代码
      // 可选：给个小提示
      outStruct.value = "// 已切换框架，请重新粘贴 SQL 生成代码";
    }
  },
  { immediate: false }
);

// script 部分
const displayedOutput = computed(() => {
  if (!outStruct.value) return "等待生成 Go struct...";
  // 保险起见再 trim 一次（防御性编程）
  return outStruct.value.trim();
});
</script>


