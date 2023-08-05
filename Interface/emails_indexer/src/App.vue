<template>
  <!--Global screen div-->
  <div class="isolate bg-gray-900 py-16 sm:py-24 lg:py-32 min-h-screen overflow-auto">

    <!--Title and input search-->
    <div class="flex m-8 p-2 px-24">
      <!--Title-->
      <div class="flex-1">
        <h2 class="text-4xl font-bold tracking-tight text-white whitespace-nowrap">Welcome to emails search!!</h2>
        <p class="mt-2 text-lg leading-8 text-gray-300">Personalize your search:</p>
      </div>
      <!--Input options-->
      <div class="flex flex-1 p-4">
          <label for="text" class="sr-only">Search term</label>
          <input id="search-term" name="term" type="text" required="" v-model="searchTerm" class="flex-auto rounded-md border-0 bg-gray-900 px-4 py-2 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6" placeholder="Enter a word or phrase to search" />
          <label for="number" class="sr-only">Select a number</label>
          <input id="number" name="selectedNumber" type="number" v-model="selectedNumber" class="flex-auto rounded-md border-0 bg-gray-900 px-4 py-2 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6" placeholder="Number of max results" />
          <label for="search-type" class="sr-only">Search type</label>
          <select id="search-type" v-model="searchType" class="flex-auto rounded-md border-0 bg-gray-900 px-4 py-2 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6">
            <option value="matchall">Match All</option>
            <option value="match">Match</option>
            <option value="matchphrase">Match Phrase</option>
            <option value="term">Term</option>
            <option value="querystring">Query String</option>
            <option value="prefix">Prefix</option>
            <option value="wildcard">Wildcard</option>
            <option value="fuzzy">Fuzzy</option>
            <option value="daterange">Date Range</option>
          </select>
          <button @click="handleSearch" type="submit" class="flex-none rounded-md bg-[#4682A9] px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-[#3a6c93] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">Search</button>
          
          <!--Index Options-->
          <div class="relative ml-8">
            <button type="button" class="flex items-center justify-center w-10 h-10 rounded-full bg-[#4682A9] hover:bg-[#3a6c93] focus:outline-none focus:ring focus:ring-indigo-200" @click="toggleDropdown">
              <svg class="w-6 h-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
              </svg>
            </button>
            <div v-if="showDropdown" class="absolute z-10 right-0 mt-2 w-40 bg-white rounded-lg shadow-lg ring-1 ring-white/10 focus:ring-2">
              <button type="button" class="block bg-gray-900 px-4 py-2 text-sm text-white hover:bg-gray-700 w-full text-left" @click="createItem">Create new index</button>
              <button type="button" class="block bg-gray-900 px-4 py-2 text-sm text-white hover:bg-gray-700 w-full text-left" @click="deleteItem">Delete index</button>
            </div>
          </div>
        </div>
    </div>

    <!--Background shape for decoration-->
    <div class="absolute left-1/2 top-0 -z-10 -translate-x-1/2 blur-3xl xl:-top-6" aria-hidden="true">
      <div class="aspect-[1155/678] w-[72.1875rem] bg-gradient-to-tr from-[#91C8E4] to-[#F6F4EB] opacity-50" style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)" />
    </div>

    <!--Emails view section-->
    <div class="max-h-[600px] m-8 flex">
      <!--Emails list on the left-->
      <div v-if="responseDataList.length > 0" class="w-full basis-1/3 bg-gray-700 m-4 rounded-2xl overflow-auto">
        <div v-for="(response, index) in responseDataList" :key="index" class="border-b-2 p-6" @click="showEmailContent(response)">
          <p class="text-white"><strong>From:</strong> {{ response._source.From }}</p>
          <p class="text-white"><strong>Subject:</strong> {{ response._source.Subject }}</p>
          <p class="text-white"><strong>Date:</strong> {{ response._source.Date }}</p>
        </div>
      </div>
      <div v-else class="w-full basis-1/3 bg-gray-700 m-4 rounded-2xl overflow-auto">
        <div class="p-4">
          <p class="text-white">No emails match found.</p>
        </div>
      </div>
      <!--Individual email visualization-->
      <div v-if="selectedEmail" class="w-full basis-2/3 bg-gray-700 m-4 p-10 rounded-2xl overflow-auto">
        <p class="text-white"><strong>From:</strong> {{ selectedEmail._source.From }}</p>
        <p class="text-white"><strong>Date:</strong> {{ selectedEmail._source.Date }}</p>
        <p class="text-white"><strong>Subject:</strong> {{ selectedEmail._source.Subject }}</p>
        <pre class="text-white whitespace-pre-wrap">{{ selectedEmail._source.Content }}</pre>
      </div>
    </div>
    
  </div>
</template>

<script>
export default{
  data() {
    return {
      searchTerm: "",
      searchType: "match",
      responseDataList: [],
      selectedNumber: 1,
      selectedEmail: null,
      showDropdown: false,
    };
  },
  methods: {
    async handleSearch() {
    try {
      const response = await fetch('http://localhost:3000/search/', {
          method: 'POST', // POST request to the server
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(
            {
              term: this.searchTerm,
              maxresults: this.selectedNumber.toString(), 
              searchtype: this.searchType,
            }),
        });

      if (response.ok) {
        const data = await response.json();
        //console.log(data);
        if (Array.isArray(data) && data.length > 0) {
          this.responseDataList = data;
          this.selectedEmail = null;
        } else {
          this.responseDataList = [];
          this.selectedEmail = null;
          throw new Error("Invalid response data format.");
        }
      } else {
        this.responseDataList = [];
        this.selectedEmail = null;
        console.error('Error:', response.status);
      }
    } catch (error) {
      console.error('Network Error:', error);
      }
    },
    showEmailContent(response) {
      this.selectedEmail = response;
    },
    toggleDropdown() {
      this.showDropdown = !this.showDropdown;
    },
    async createItem() {
      // Code to handle the "Create" button click
      // Add your logic here
    },
    async deleteItem() {
      // Code to handle the "Delete" button click
      // Add your logic here
    },
  },
};
</script>