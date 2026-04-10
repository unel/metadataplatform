# Доступные формы

```html
<form method="POST" action="/entities">
  <!-- явный label через for/id -->
  <div>
    <label for="name">Название <span aria-hidden="true">*</span></label>
    <input
      id="name"
      name="name"
      type="text"
      required
      aria-required="true"
      aria-describedby="name-error"
    />
    <span id="name-error" role="alert">
      <!-- ошибка валидации -->
    </span>
  </div>

  <!-- группа связанных полей -->
  <fieldset>
    <legend>Тип файла</legend>
    <label><input type="radio" name="type" value="video" /> Видео</label>
    <label><input type="radio" name="type" value="image" /> Изображение</label>
  </fieldset>

  <button type="submit">Создать</button>
</form>
```

## Placeholder — не замена label

```html
<!-- плохо — placeholder исчезает при вводе -->
<input type="email" placeholder="Введите email" />

<!-- хорошо -->
<label for="email">Email</label>
<input id="email" type="email" placeholder="example@mail.com" />
```
