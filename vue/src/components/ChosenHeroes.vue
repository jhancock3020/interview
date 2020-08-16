<template>     <!-- HTML -->
  <div>
    <select v-model="chosenHero"><!-- Create a drop-down list with four options  -->
    <!-- Create a drop-down list with four options  -->
    <!-- You can use the v-model directive to create two-way data bindings on form input, textarea, and select elements. -->
    <!-- Two-way binding means that any data-related changes affecting the model are immediately 
     propagated to the matching view(s), and that any changes made in the view(s) (say, by the user) 
     are immediately reflected in the underlying model. When app data changes, so does the UI, and conversely. -->
     <!-- placeholder value -->
      <option :value="null">Select a hero</option>

      <!-- available heroes -->
      <!-- We can set Vue to track each element using a key. 
      This would cause it to move elements rather than replacing values. -->
      <option v-for="hero in anotherLabel"
              :key="hero.name" 
              :value="hero.name">
        {{ hero.name }} <!-- display hero name -->
      </option>
    </select>
    <span>&nbsp;</span>
    <!-- A commonly used entity in HTML is the non-breaking space: &nbsp;
    A non-breaking space is a space that will not break into a new line. 
    The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent anything.-->
    <button @click="addHero(chosenHero)"
            :disabled="chosenHero === null || chosenHeroes.length === 3">Add Hero</button>
    <span> &nbsp;</span>   
    <button @click="alertFunct()">Launch Mission</button>    
    <br>
    <h3>Chosen Heroes</h3>
    <div class="chosen-heroes">
      <div v-for="(hero, i) in chosenHeroes"
           :key="hero.name">
        <strong>Slot {{ i + 1 }}:</strong>
        <!-- The <strong> tag is used to define text 
        with strong importance. The content inside is typically displayed in bold. -->
        <Hero :hero="hero"
              @removeHero="removeHero(hero)" />
      </div>
    </div>
  </div>
</template>

<script> //Javascript
import Hero from "./Hero";

export default {
  components: {
    Hero
  },
  props: ["heroes"],
  data() {
    return {
      chosenHero: null,
      chosenHeroes: [],
      anotherLabel: this.heroes,
      testHero: 0,
    };
  },
  methods: {
    addHero(name) {
      //this.name.isChosen = true;
      this.chosenHeroes.push({ name });
      this.chosenHero = null;
      this.anotherLabel = this.anotherLabel.filter(h => h.name != name);
    },
    removeHero(hero) {
      hero.isChosen = true;
      this.chosenHeroes = this.chosenHeroes.filter(h => h.name != hero.name);
      this.anotherLabel.push(hero);
    },
    alertFunct() {
      if(this.chosenHeroes.length == 3){
        alert("Mission complete");
      }
      else{
        alert("We need three heroes");
      }
    }
  }
};
</script>

<style scoped>
.chosen-heroes {
  display: flex; 
  flex-flow: column;
  align-items: center;
}
</style>


