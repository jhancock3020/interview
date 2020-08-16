<template>
  <div>
    <!-- justice leage application begins here -->
    <h1 id="jl">Justice League Mission Planner</h1>

    <ul class="roster">
      <h3>Roster:</h3>
      <li v-for="hero in heroes"
          :key="hero.name">
        <!-- to do: conditionally display this span -->
        <span v-if = "hero.isChosen">✔ &nbsp;</span>

        <span>{{ hero.name }}&nbsp;</span>
        <span class="edit"
              @click="editHero(hero)">edit</span>
      </li>
      <br>
      <input type="text"
             placeholder="new name"
             v-model="newName"
             v-if="isEdit"
             @keyup.enter="changeName"
             @blur="clear">
      <br>
      <span v-if="isEdit">enter to submit, click outside the box to cancel</span>
    </ul>
    <ChosenHeroes :heroes="heroes" />
  </div>
</template>

<script>
import ChosenHeroes from "./components/ChosenHeroes.vue";
export default {
  components: {
    ChosenHeroes
  },
  
  data() {
    return {
      heroes: [
        { name: "Superman", isChosen: false },
        { name: "Batman" , isChosen: false},
        { name: "Aquaman" , isChosen: false},
        { name: "Wonder Woman" , isChosen: false},
        { name: "Green Lantern", isChosen: false },
        { name: "Martian Manhunter" , isChosen: false},
        { name: "Flash", isChosen: false }
      ],
      newName: "",
      isEdit: false,
      heroToModify: null,
    };
  },
  methods: {
    editHero(hero) {
      this.isEdit = true;
      this.newName = hero.name;
      this.heroToModify = hero;
      hero.isChosen = true;
    },

    changeName() {
      this.heroToModify.name = this.newName;
      this.isEdit = false;
    },

    clear() {
      this.heroToModify = null;
      this.newName = "";
      this.isEdit = false;
    },
  }
};
</script>

<style scoped>
ul.roster {
  text-align: left;
}
.edit {
  color: blue;
  text-decoration: underline;
}
.edit:hover {
  cursor: pointer;
}
</style>
