<script>
export default {
    props: {
        isMine: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            isOpen: false
        }
    },
    mounted() {
        document.addEventListener('click', this.closeMenu);
    },
    beforeUnmount() {
        document.removeEventListener('click', this.closeMenu);
    },
    methods: {
        toggleMenu() {
            this.isOpen = !this.isOpen;
        },
        closeMenu() {
            this.isOpen = false;
        },
        selectOption(action) {
            this.$emit(action);
            this.closeMenu();
        }
    }
}
</script>

<template>
    <div class="message-menu">
        <button class="menu-button" @click.stop="toggleMenu">⋮</button>
        
        <div v-if="isOpen" class="dropdown-content">
            <div class="dropdown-item" @click.stop="selectOption('reply')">Reply</div>
            <div class="dropdown-item" @click.stop="selectOption('forward')">Forward</div>
            <div class="dropdown-item" @click.stop="selectOption('react')">React</div>
            <div v-if="isMine" class="dropdown-item delete" @click.stop="selectOption('delete')">Delete</div>
        </div>
    </div>
</template>

<style scoped>
.message-menu {
    position: relative;
    display: flex;
    align-items: center;
}

.menu-button {
    background: none;
    border: none;
    font-size: 20px;
    font-weight: bold;
    cursor: pointer;
    color: #888;
    border-radius: 50%;
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.2s;
}

.menu-button:hover {
    background-color: rgba(0, 0, 0, 0.05);
}

.dropdown-content {
    position: absolute;
    top: 35px;
    right: 0;
    background-color: var(--color-2-a);
    min-width: 120px;
    box-shadow: 0px 4px 12px rgba(0,0,0,0.15);
    border-radius: 8px;
    z-index: 100;
    overflow: hidden;
    border: 1px solid #eaeaea;
}

.dropdown-item {
    color: var(--text-color-2);
    padding: 10px 16px;
    display: block;
    cursor: pointer;
    font-size: 14px;
    transition: background-color 0.2s;
}

.dropdown-item:hover {
    background-color: var(--color-5);
}

.dropdown-item.delete {
    color: #d9534f;
}
</style>