<template> <!-- Use <template> to hold some content that will be hidden when the page loads. Use JavaScript to display it:-->
    <div class="card"><!-- A <div> section in a document that is styled with CSS -->
        <h3 class="close" @click="$emit('removeHero')">&times;</h3><!--use the @click directive to listen to DOM events 
        and run some JavaScript when they’re triggered. 
        $emit() function that allows you to pass custom events up the component tree 
        &times; displays x-->
        <h1>{{ hero.name }}</h1><!-- h1 displays hero name -->
    </div>
</template>

<script>
export default {// Default exports 
//are useful to export only a single object, function, variable. During the import, we can use any name to import.
  props: ["hero"]//Props are custom attributes you can register on a component. 
  //When a value is passed to a prop attribute, it becomes a property on that component instance. 
};
</script>

<style scoped>
.card {
  border: 1px solid black;
  margin-top: 1em;
  margin-bottom: 1em;
  width: 500px;
}
.card h3 {
    text-align: end;
    margin-right: 1em;
}
.card h3:hover {
    cursor: pointer;
}
</style>
