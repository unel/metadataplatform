# Поиск исходящих extends по коду

## Go
- embedding структур из другого пакета (`type Foo struct { Bar }`)
- реализация интерфейса определённого в другом пакете
- функции-обёртки которые вызывают базовую функцию и добавляют логику

## SvelteKit

### lib/components/
- компонент принимает другой компонент через slot и оборачивает его
- `extends` в props типе (`interface Props extends BaseProps`)

### lib/stores/
- `derived()` стор поверх базового стора с дополнительной логикой

### lib/types/
- `extends` в TypeScript интерфейсе
- `type Foo = Bar & { ... }`

## CSS
- `@extend` или переопределение CSS custom properties базового компонента
- класс который наследует стили базового через cascade
