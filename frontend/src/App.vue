<script>
import ImportWishlist from "./components/ImportWishlist.vue";
import UserInfo from "./components/UserInfo.vue";
import Wishlist from "./components/Wishlist.vue";

export default {
    name: "App",
    components: {
        ImportWishlist,
        UserInfo,
        Wishlist,
    },
    data() {
        return {
            userId: null,
        };
    },
    methods: {
        updateUserId(userId) {
            this.userId = userId;
        } 
    },
};
</script>

<template>
    <header>
        <h1>Present-Roulette</h1>
        <UserInfo v-bind:userId="userId" @receivedUserId="updateUserId" />
    </header>

    <main>
        <div v-if="userId">
            <ImportWishlist />
            <Wishlist />
        </div>
        <div v-else>
            <p>You are not logged in</p>
        </div>
    </main>
</template>

<style></style>
