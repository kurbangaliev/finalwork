--
-- PostgreSQL database dump
--

-- Dumped from database version 18.1 (Debian 18.1-1.pgdg13+2)
-- Dumped by pg_dump version 18.1

-- Started on 2026-01-27 18:18:29 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 223 (class 1259 OID 24597)
-- Name: managers; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.managers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    job_title character varying(255),
    address character varying(255),
    phone character varying(255),
    email character varying(255),
    schedule character varying(255),
    image character varying(255)
);


ALTER TABLE public.managers OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 24596)
-- Name: managers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.managers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.managers_id_seq OWNER TO postgres;

--
-- TOC entry 3482 (class 0 OID 0)
-- Dependencies: 222
-- Name: managers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.managers_id_seq OWNED BY public.managers.id;


--
-- TOC entry 221 (class 1259 OID 16486)
-- Name: news; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.news (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    content text,
    date character varying(20),
    image character varying(1024)
);


ALTER TABLE public.news OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16485)
-- Name: news_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.news_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.news_id_seq OWNER TO postgres;

--
-- TOC entry 3483 (class 0 OID 0)
-- Dependencies: 220
-- Name: news_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.news_id_seq OWNED BY public.news.id;


--
-- TOC entry 225 (class 1259 OID 24608)
-- Name: partners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.partners (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255),
    image character varying(255)
);


ALTER TABLE public.partners OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 24607)
-- Name: partners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.partners_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.partners_id_seq OWNER TO postgres;

--
-- TOC entry 3484 (class 0 OID 0)
-- Dependencies: 224
-- Name: partners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.partners_id_seq OWNED BY public.partners.id;


--
-- TOC entry 219 (class 1259 OID 16386)
-- Name: slogans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.slogans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 2147483647
    CACHE 1;


ALTER SEQUENCE public.slogans_id_seq OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 32789)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    login character varying(255),
    password character varying(255),
    role character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 32788)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3485 (class 0 OID 0)
-- Dependencies: 226
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3306 (class 2604 OID 24600)
-- Name: managers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.managers ALTER COLUMN id SET DEFAULT nextval('public.managers_id_seq'::regclass);


--
-- TOC entry 3305 (class 2604 OID 16489)
-- Name: news id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news ALTER COLUMN id SET DEFAULT nextval('public.news_id_seq'::regclass);


--
-- TOC entry 3307 (class 2604 OID 24611)
-- Name: partners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners ALTER COLUMN id SET DEFAULT nextval('public.partners_id_seq'::regclass);


--
-- TOC entry 3308 (class 2604 OID 32792)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3472 (class 0 OID 24597)
-- Dependencies: 223
-- Data for Name: managers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.managers (id, created_at, updated_at, deleted_at, name, job_title, address, phone, email, schedule, image) FROM stdin;
1	2026-01-24 18:05:06.075983+00	2026-01-24 19:15:26.154439+00	\N	ДЮСЕНОВ Ринат Тасбулатович	Генеральный директор ТОО «DAR RAIL»	Республика Казахстан, г. Астана, проспект Ракымжан Кошкарбаев, 1/5, 6-этаж	+7 7172 39 99 88 (вн.510)	info@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074412130200_Dyussenov.jpg
2	2026-01-24 19:17:33.301472+00	2026-01-24 19:17:33.301472+00	\N	СМИРНОВА Лариса Иосифовна	Директор Филиала	Республика Казахстан, г. Алматы, проспект Достык, 291/23, 2-этаж	+7 7172 39 99 88	info-almaty@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074367376000_Smirnova.jpg
3	2026-01-24 19:21:17.899034+00	2026-01-24 19:21:17.899034+00	\N	ЕРМОЛЬЧЕВ Александр Александрович	Директор регионального подразделения на ст. Тобол	Республика Казахстан, пос.Тобол , улица Станционная, 1, 1-этаж	+7 714 64 00 26 (вн.5101)	Info-tobol@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074379542800_Ermolchev.jpg
4	2026-01-24 19:22:45.490981+00	2026-01-24 19:22:45.490981+00	\N	Байымбетов Арман Санатарович	Директор регионального подразделения на ст. Павлодар	Республика Казахстан, город Павлодар, улица Путейская, 2, 3	+7 7182 37 22 84 (вн.521)	Info-ekibastuz@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074428197900_Bayimbetov.jpg
5	2026-01-24 19:24:16.740866+00	2026-01-24 19:24:16.740866+00	\N	АСКАРОВ Марат Сабырович	Директор регионального подразделения на ст. Нур-Султан	Республика Казахстан, г. Астана, улица Шынтас, 8, 3-этаж	+7 7172 93 43 53 (вн.570)	info@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074328683900_Askarov.jpg
6	2026-01-24 19:25:12.583895+00	2026-01-24 19:25:12.583895+00	\N	БЕКЕНОВ Бестыбай Каратаевич	Начальник участка на ст. Екибастуз	Республика Казахстан, г. Экибастуз, улица Астана, 31, 3-этаж	+7 7187 22 72 35 (вн.520)	Info-ekibastuz@darrail.com	09:00 – 18:00	http://localhost:8081/uploads/managers/1769097074397199900_Bekenov.jpg
\.


--
-- TOC entry 3470 (class 0 OID 16486)
-- Dependencies: 221
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.news (id, created_at, updated_at, deleted_at, title, content, date, image) FROM stdin;
3	2026-01-21 17:03:52.802232+00	2026-01-21 17:27:33.036073+00	\N	ОО "Локальный профессиональный союз DAR Rail" подарил детям наших сотрудников настоящее летнее приключение в лагерях	ОО "Локальный профессиональный союз DAR Rail" подарил детям наших сотрудников настоящее летнее приключение в лагерях\n\n \n\nЭтим летом дети сотрудников РП Тобол, Павлодар, участков Экибастуз и Аксу провели незабываемое время в  летних лагерях «Чайка» и «Жас  Даурен». Дети были обеспечены ярким и активным отдыхом и укрепили здоровье, а также развили творческие и спортивные способности и получили массу положительных эмоций.\n\n \n\nОба лагеря располагаются в живописных местах, вдали от городского шума, среди соснового леса. Дети жили в уютных корпусах, окружённых зеленью, с удобствами и тёплой атмосферой. Каждый день был насыщен различными мероприятиями — утренние зарядки, спортивные соревнования, кружки по интересам, театральные постановки, мастер-классы и квесты.\n\n \n\nОсобой популярностью пользовались занятия по рисованию, робототехнике и танцам. Ребята принимали участие в конкурсах талантов, творческих вечерах и интеллектуальных играх. Не обошлось и без походов, купания в реке под присмотром вожатых и вечерних костров с песнями под гитару. Питание в лагере было сбалансированным и вкусным, что особенно порадовало родителей. За здоровьем детей следили квалифицированные медики, а вожатые создавали тёплую и дружелюбную атмосферу.\n\n \n\nПо возвращении домой дети делились впечатлениями, рассказывали о новых друзьях, победах и открытиях. Многие уже выразили желание поехать в лагерь на следующий год. Сейчас в летнем лагере «Звездный Бурабай» на озере Катарколь отдыхают детки работников РП Астана, участков Есиль и Жана-Аул. Надеемся что им тоже понравится время проведенное в лагере.\n\n \n\nТакие поездки — это не только вклад в здоровье и развитие ребёнка, но и большая поддержка для сотрудников, которые уверены, что их дети проводят время с пользой и в безопасности. Коллектив Профсоюзной организации выражает руководству Компании  искреннюю благодарность за организацию такого важного и нужного мероприятия	2025-01-31	http://localhost:8081/uploads/news/1769014996647459200_n4.png
5	2026-01-21 18:11:47.275308+00	2026-01-21 18:11:47.275308+00	\N	Встреча машинистов DAR RAIL со студентами железнодорожного колледжа	В рамках мероприятий Года рабочих профессий машинисты ТОО «DAR RAIL» посетили столичный железнодорожный колледж, где провели открытую встречу со студентами третьего курса специальности «Эксплуатация, ремонт и техническое обслуживание тягового подвижного состава».\n\nВстречу открыл заместитель директора колледжа по учебно-производственной практике Бауыржан Ахметов. В ходе визита машинисты электровоза Иван Ильдеряков и Әбілқайыр Дүйсебек из регионального подразделения станции Астана рассказали о своей работе, поделились практическими случаями и ответили на многочисленные вопросы студентов. Также участникам мероприятия был показан видеоролик о профессии.\n\nПодобные встречи играют важную роль в профориентации: они помогают молодым людям лучше понять специфику рабочих профессий и сделать осознанный выбор профессионального пути. DAR RAIL последовательно поддерживает такие инициативы, ведь развитие кадрового потенциала — один из ключевых приоритетов компании.\n\n \n\nКомпания выражает благодарность преподавателям и студентам колледжа за тёплый приём и активное участие. Уверены, что среди слушателей — будущие профессионалы, которые в будущем будут с честью работать на благо казахстанской железнодорожной отрасли.	2025-04-24T23:11:47	http://localhost:8081/uploads/news/1769008779944003000_n1.jpg
4	2026-01-21 17:07:42.730261+00	2026-01-21 17:07:42.730261+00	\N	Встреча машинистов DAR RAIL со студентами железнодорожного колледжа	В рамках мероприятий Года рабочих профессий машинисты ТОО «DAR RAIL» посетили столичный железнодорожный колледж, где провели открытую встречу со студентами третьего курса специальности «Эксплуатация, ремонт и техническое обслуживание тягового подвижного состава».\n\nВстречу открыл заместитель директора колледжа по учебно-производственной практике Бауыржан Ахметов. В ходе визита машинисты электровоза Иван Ильдеряков и Әбілқайыр Дүйсебек из регионального подразделения станции Астана рассказали о своей работе, поделились практическими случаями и ответили на многочисленные вопросы студентов. Также участникам мероприятия был показан видеоролик о профессии.\n\nПодобные встречи играют важную роль в профориентации: они помогают молодым людям лучше понять специфику рабочих профессий и сделать осознанный выбор профессионального пути. DAR RAIL последовательно поддерживает такие инициативы, ведь развитие кадрового потенциала — один из ключевых приоритетов компании.\n\n \n\nКомпания выражает благодарность преподавателям и студентам колледжа за тёплый приём и активное участие. Уверены, что среди слушателей — будущие профессионалы, которые в будущем будут с честью работать на благо казахстанской железнодорожной отрасли.	2025-04-24T22:07:42	http://localhost:8081/uploads/news/1769015232138182300_n5.png
1	2026-01-21 15:34:53.352491+00	2026-01-21 15:34:53.352491+00	\N	DAR Rail отпраздновал День железнодорожника и отметил лучших сотрудников	В канун профессионального праздника — Дня работников транспорта — в центральном офисе DAR Rail прошло торжественное собрание с онлайн-подключением региональных подразделений компании: Астаны, Павлодара, Тобола и Экибастуза.\n\nОперационный директор Асхат Асылханович Тусупканов обратился к коллегам с поздравлением, подчеркнув важность этого дня для всех, кто связал жизнь с железной дорогой:\n\n«Железная дорога — это не просто работа. Это образ жизни, требующий точности, выдержки и настоящего командного духа. Пусть каждый рабочий день будет безопасным и спокойным, каждый рейс — успешным, а жизнь наполнена теплом, уютом и радостными событиями».\n\nНа встрече подвели итоги года, отметили ключевые достижения компании и выразили благодарность сотрудникам, которые внесли значительный вклад в развитие DAR Rail.	2025-08-01T20:34:53	http://localhost:8081/uploads/news/1769009319880120000_n2.png
6	2026-01-22 21:00:32.719018+00	2026-01-22 21:00:32.719018+00	\N	DAR Rail отпраздновал День железнодорожника и отметил лучших сотрудников	В канун профессионального праздника — Дня работников транспорта — в центральном офисе DAR Rail прошло торжественное собрание с онлайн-подключением региональных подразделений компании: Астаны, Павлодара, Тобола и Экибастуза.\n\nОперационный директор Асхат Асылханович Тусупканов обратился к коллегам с поздравлением, подчеркнув важность этого дня для всех, кто связал жизнь с железной дорогой:\n\n«Железная дорога — это не просто работа. Это образ жизни, требующий точности, выдержки и настоящего командного духа. Пусть каждый рабочий день будет безопасным и спокойным, каждый рейс — успешным, а жизнь наполнена теплом, уютом и радостными событиями».\n\nНа встрече подвели итоги года, отметили ключевые достижения компании и выразили благодарность сотрудникам, которые внесли значительный вклад в развитие DAR Rail.\nОтраслевые и внутренние награды\n\nПо решению наградной комиссии DAR Rail, возглавляемой генеральным директором Ринатом Дюсеновым, сотрудники компании были удостоены ведомственных и корпоративных наград.\n\n    Почётная грамота Министра транспорта РК —\n     Сергей Борушевский, директор по финансам;\n     Азамат Шарсеитов, директор по развитию бизнеса и обеспечению.\n\n    Благодарственное письмо Министра транспорта РК —\n     Сергей Деринг, руководитель службы по операционной деятельности;	2025-08-01T02:00:32	http://localhost:8081/uploads/news/1769009319880120000_n2.png
2	2026-01-21 15:36:46.46187+00	2026-01-24 18:38:21.647443+00	\N	V Евразийский форум по безопасности и цифровизации на транспорте	4 октября 2024 года в Алматы, на базе ALT университета имени Мухамеджана Тынышпаева, прошло важное отраслевое событие — V Евразийский форум по безопасности и цифровизации на транспорте, в рамках которого состоялась выставка компаний. В мероприятии приняли участие представители КТЖ, РЖД, производителей железнодорожной техники и оборудования, а также поставщики ИТ-решений. \n\n \n\nКомпания DAR RAIL активно участвовала в мероприятии, где генеральный директор Ринат Дюсенов представил инновационную систему АСУП (Автоматизированная система управления перевозками) в рамках панельной сессии, посвященной цифровой трансформации железнодорожной отрасли. Эта система значительно улучшает эффективность управления, оптимизируя процессы и снижая затраты. Презентация вызвала живой интерес у участников, подчеркивая важность внедрения цифровых решений в современном железнодорожном транспорте.\n\n \n\nКроме того, компания была представлена на выставочном стенде Ассоциации казахстанских грузовых железнодорожных перевозчиков. Там посетители могли ознакомиться с информацией о деятельности DAR RAIL и внедряемых ИТ-решениях, представленными в наглядной форме. \n\n \n\nСо стендом DAR RAIL ознакомились: заместитель председателя правления КТЖ Талгат Алдыбергенов, директор департамента безопасности движения КТЖ Марат Шакенов и заместитель генерального директора РЖД — директор департамента безопасности движения Шевкет Шайдуллин.	2024-10-08	http://localhost:8081/uploads/news/1769009789818937500_n3.jpg
\.


--
-- TOC entry 3474 (class 0 OID 24608)
-- Dependencies: 225
-- Data for Name: partners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.partners (id, created_at, updated_at, deleted_at, name, image) FROM stdin;
\.


--
-- TOC entry 3476 (class 0 OID 32789)
-- Dependencies: 227
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, updated_at, deleted_at, login, password, role) FROM stdin;
2	2026-01-27 18:04:22.446353+00	2026-01-27 18:04:22.446353+00	\N	admin	ISMvKXpXpadDiUoOSoAfww==	admin
\.


--
-- TOC entry 3486 (class 0 OID 0)
-- Dependencies: 222
-- Name: managers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.managers_id_seq', 6, true);


--
-- TOC entry 3487 (class 0 OID 0)
-- Dependencies: 220
-- Name: news_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.news_id_seq', 6, true);


--
-- TOC entry 3488 (class 0 OID 0)
-- Dependencies: 224
-- Name: partners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.partners_id_seq', 1, false);


--
-- TOC entry 3489 (class 0 OID 0)
-- Dependencies: 219
-- Name: slogans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.slogans_id_seq', 3, true);


--
-- TOC entry 3490 (class 0 OID 0)
-- Dependencies: 226
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 3314 (class 2606 OID 24605)
-- Name: managers managers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.managers
    ADD CONSTRAINT managers_pkey PRIMARY KEY (id);


--
-- TOC entry 3311 (class 2606 OID 16494)
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- TOC entry 3317 (class 2606 OID 24616)
-- Name: partners partners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners
    ADD CONSTRAINT partners_pkey PRIMARY KEY (id);


--
-- TOC entry 3320 (class 2606 OID 32797)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3312 (class 1259 OID 24606)
-- Name: idx_managers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_managers_deleted_at ON public.managers USING btree (deleted_at);


--
-- TOC entry 3309 (class 1259 OID 16495)
-- Name: idx_news_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_news_deleted_at ON public.news USING btree (deleted_at);


--
-- TOC entry 3315 (class 1259 OID 24617)
-- Name: idx_partners_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_partners_deleted_at ON public.partners USING btree (deleted_at);


--
-- TOC entry 3318 (class 1259 OID 32798)
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


-- Completed on 2026-01-27 18:18:30 UTC

--
-- PostgreSQL database dump complete
--


