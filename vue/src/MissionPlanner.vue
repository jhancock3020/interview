<template>
  <div>
    <!-- justice leage application begins here -->
    <h1 id="jl">Justice League Mission Planner</h1>

    <ul class="roster">
      <h3>Roster:</h3>
      <li v-for="hero in heroes"
          :key="hero.name">
        <!-- If hero.chosen in heroes array is true (ie hero is in heroesChosen list), show checkmark -->
        <span v-if = "hero.chosen"> ✔ &nbsp;</span>

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
    <ChosenHeroes :heroes="heroes" 
    @checkmark="handler"/><!-- From this instance of ChosenHeroes to MissionPlanner's handler method -->
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
      heroes: [ // name is currently shown name, 
      //chosen keeps track if heroes elem is in chosenHeroes array and if checkmark should be shown as a result, 
      //prevName is prevous hero.name
        { name: "Superman",chosen: false, prevName: "Superman"},
        { name: "Batman",chosen: false, prevName:  "Batman"},
        { name: "Aquaman",chosen: false, prevName:  "Aquaman"},
        { name: "Wonder Woman",chosen: false, prevName: "Wonder Woman"},
        { name: "Green Lantern",chosen: false, prevName:"Green Lantern"},
        { name: "Martian Manhunter",chosen: false, prevName:"Martian Manhunter" },
        { name: "Flash",chosen: false, prevName: "Flash"}
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
    handler(value, bool){
      var i;// Goes through heroes using parameters from ChosenHeroes instance, 
      //if hero.name is equal to chosenHero name AND bool is true, hero is in array
      //and gets checkmark, if hero.name is equal to chosenHero name AND boolean is false
      //hero is no longer in chosenHeroes array and loses checkmark
      for (i = 0; i < this.heroes.length; i++) {
        if(this.heroes[i].name == value && bool){
          this.heroes[i].chosen = true;
        }else if(this.heroes[i].name == value && bool == false){
          this.heroes[i].chosen = false;
        }
      }
    }
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
