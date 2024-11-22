CREATE TABLE db.events (
    code String NOT NULL,
    gender Enum('male' = 1, 'female' = 2) NOT NULL,
    sport String NOT NULL,
    additional_info String,
    country String NOT NULL,
    region String NOT NULL,
    locality String NOT NULL,
    n_participants UInt32 NOT NULL,
    stage String NOT NULL,
    start_date Datetime NOT NULL,
    end_date Datetime NOT NULL,
)
ENGINE = MergeTree()
ORDER BY (gender, start_date, code);

CREATE TABLE db.event_restrictions (
    code String NOT NULL,
    gender Enum('male' = 1, 'female' = 2) NOT NULL,
    left_bound UInt32 NOT NULL,
    right_bound UInt32 NOT NULL,
    extra_mapping String,
)
ENGINE = MergeTree()
ORDER BY (left_bound, right_bound, code);

INSERT INTO db.events (code, gender, sport, additional_info, country, region, locality, n_participants, stage, start_date, end_date)
VALUES
    ('100001', 'male', 'Прыжки через гаражи', 'Прыжки через гаражи, до 70 кг', 'Россия', 'Московская область', 'г. Москва', 150, 'Чемпионат двора', '2024-12-01 12:00:00', '2024-12-02 18:00:00'),
    ('100002', 'female', 'Бег по раскалённым углям', 'На скорость и грацию', 'Россия', 'Ленинградская область', 'г. Санкт-Петербург', 75, 'Междворовой турнир', '2024-12-15 10:00:00', '2024-12-15 15:00:00'),
    ('100003', 'male', 'Катание на табуретках', 'С препятствиями, до 90 кг', 'Россия', 'Тюменская область', 'г. Тюмень', 120, 'Кубок Тюмени', '2025-01-10 09:00:00', '2025-01-12 17:00:00'),
    ('100004', 'female', 'Поднятие самоваров', 'Слева и справа', 'Россия', 'Тульская область', 'г. Тула', 60, 'Самоварный баттл', '2025-02-20 11:00:00', '2025-02-21 13:00:00'),
    ('100005', 'male', 'Борьба на качелях', 'С прыжками и подсечками', 'Россия', 'Свердловская область', 'г. Екатеринбург', 85, 'Финал России', '2025-03-05 14:00:00', '2025-03-06 16:00:00'),
    ('100006', 'female', 'Перетягивание кота', 'На выносливость', 'Россия', 'Курская область', 'г. Курск', 50, 'Дворовая лига', '2025-04-01 08:00:00', '2025-04-01 12:00:00'),
    ('100007', 'male', 'Бег с утюгами', 'На скорость и чистоту', 'Россия', 'Самарская область', 'г. Самара', 140, 'Кубок Утюга', '2025-05-10 07:00:00', '2025-05-11 20:00:00'),
    ('100008', 'female', 'Гонки на метлах', 'С фигурами пилотажа', 'Россия', 'Рязанская область', 'г. Рязань', 95, 'Межрегиональные старты', '2025-06-15 09:30:00', '2025-06-15 14:30:00'),
    ('100009', 'male', 'Толкание самосвала', 'В горку, до 100 кг', 'Россия', 'Челябинская область', 'г. Челябинск', 200, 'Железный чемпионат', '2025-07-25 10:00:00', '2025-07-27 18:00:00'),
    ('100010', 'female', 'Бег на каблуках', '100 метровка', 'Россия', 'Краснодарский край', 'г. Краснодар', 130, 'Чемпионат Юга', '2025-08-20 16:00:00', '2025-08-20 19:00:00'),
    ('100011', 'male', 'Метание тапков', 'На дальность', 'Россия', 'Ростовская область', 'г. Ростов-на-Дону', 180, 'Кубок тапка', '2025-09-10 12:00:00', '2025-09-10 16:00:00'),
    ('100012', 'female', 'Соревнования по глажке', 'На скорость и аккуратность', 'Россия', 'Пензенская область', 'г. Пенза', 100, 'Гладильный турнир', '2025-10-05 08:00:00', '2025-10-06 14:00:00'),
    ('100013', 'male', 'Перетягивание каната на велосипедах', 'С элементами акробатики', 'Россия', 'Иркутская область', 'г. Иркутск', 110, 'Кубок Байкала', '2025-11-15 13:00:00', '2025-11-16 17:00:00'),
    ('100014', 'female', 'Марафон по скандинавской ходьбе', 'По песку и болоту', 'Россия', 'Калининградская область', 'г. Калининград', 65, 'Марафон Балтики', '2025-12-10 07:00:00', '2025-12-10 21:00:00'),
    ('100015', 'male', 'Подводное шахматное дерби', 'Шахматы в акваланге', 'Россия', 'Приморский край', 'г. Владивосток', 40, 'Чемпионат Приморья', '2025-12-20 10:00:00', '2025-12-21 18:00:00'),
    ('100016', 'female', 'Скалолазание на шкафы', 'На скорость', 'Россия', 'Татарстан', 'г. Казань', 90, 'Кубок Казани', '2026-01-15 09:00:00', '2026-01-15 17:00:00');


INSERT INTO db.event_restrictions (code, gender, left_bound, right_bound, extra_mapping)
VALUES
    ('0321412542', 'male', 18, 80, 'мужчина'),
    ('100001', 'male', 16, 40, 'юноша'),
    ('100001', 'male', 41, 60, 'мужчина'),
    ('100002', 'female', 18, 35, 'девушка'),
    ('100002', 'female', 36, 50, 'женщина'),
    ('100003', 'male', 20, 35, 'юноша'),
    ('100003', 'male', 36, 90, 'мужчина'),
    ('100004', 'female', 18, 30, 'девушка'),
    ('100004', 'female', 31, 50, 'женщина'),
    ('100005', 'male', 22, 40, 'юноша'),
    ('100005', 'male', 41, 80, 'мужчина'),
    ('100006', 'female', 15, 20, 'юниорка'),
    ('100006', 'female', 21, 30, 'девушка'),
    ('100007', 'male', 19, 30, 'юноша'),
    ('100007', 'male', 31, 45, 'мужчина'),
    ('100008', 'female', 18, 25, 'девушка'),
    ('100008', 'female', 26, 60, 'женщина'),
    ('100009', 'male', 25, 50, 'мужчина'),
    ('100009', 'male', 51, 100, 'мужчина'),
    ('100010', 'female', 21, 30, 'девушка'),
    ('100010', 'female', 31, 40, 'женщина'),
    ('100011', 'male', 16, 30, 'юноша'),
    ('100011', 'male', 31, 45, 'мужчина'),
    ('100012', 'female', 18, 30, 'девушка'),
    ('100012', 'female', 31, 50, 'женщина'),
    ('100013', 'male', 20, 40, 'юноша'),
    ('100013', 'male', 41, 60, 'мужчина'),
    ('100014', 'female', 25, 40, 'женщина'),
    ('100014', 'female', 41, 55, 'женщина'),
    ('100015', 'male', 18, 30, 'юноша'),
    ('100015', 'male', 31, 45, 'мужчина'),
    ('100016', 'female', 20, 40, 'девушка'),
    ('100016', 'female', 41, 60, 'женщина');