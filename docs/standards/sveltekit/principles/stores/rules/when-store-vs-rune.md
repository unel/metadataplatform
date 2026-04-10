# Когда store, когда rune

## Используй rune ($state) когда

- Состояние локально для компонента или дерева компонентов
- Состояние передаётся через props
- Используется паттерн createXxx с rune-функцией

## Используй store когда

- Состояние нужно разделять между несвязанными компонентами
- Состояние живёт вне компонентного дерева (auth сессия, глобальные уведомления)
- Нужна подписка из не-svelte кода

```typescript
// lib/stores/auth.ts
import { writable, derived } from 'svelte/store'

interface AuthState {
  user: User | null
  loading: boolean
}

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthState>({
    user: null,
    loading: false
  })

  return {
    subscribe,
    login: async (credentials: Credentials) => {
      update(s => ({ ...s, loading: true }))
      const user = await api.login(credentials)
      set({ user, loading: false })
    },
    logout: () => set({ user: null, loading: false })
  }
}

export const auth = createAuthStore()
export const isAuthenticated = derived(auth, $auth => $auth.user !== null)
```
