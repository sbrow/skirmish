--
-- PostgreSQL database dump
--

-- Dumped from database version 10.2
-- Dumped by pg_dump version 10.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

DROP RULE IF EXISTS wisp_insert_default ON public.wisp;
DROP RULE IF EXISTS vi_insert_default ON public.vi;
DROP RULE IF EXISTS tinsel_insert_default ON public.tinsel;
DROP RULE IF EXISTS tendril_insert_default ON public.tendril;
DROP RULE IF EXISTS scuttler_insert_default ON public.scuttler;
DROP RULE IF EXISTS scinter_insert_default ON public.scinter;
DROP RULE IF EXISTS ravat_insert_default ON public.ravat;
DROP RULE IF EXISTS lilith_insert_default ON public.lilith;
DROP RULE IF EXISTS igrath_insert_default ON public.igrath;
DROP RULE IF EXISTS bast_insert_default ON public.bast;
ALTER TABLE IF EXISTS ONLY public.zones DROP CONSTRAINT IF EXISTS zones_pkey;
ALTER TABLE IF EXISTS ONLY public.wisp DROP CONSTRAINT IF EXISTS wisp_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_heroes DROP CONSTRAINT IF EXISTS wisp_heroes_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_events DROP CONSTRAINT IF EXISTS wisp_events_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_continuous_events DROP CONSTRAINT IF EXISTS wisp_continuous_events_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_channeled_heroes DROP CONSTRAINT IF EXISTS wisp_channeled_heroes_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_channeled_events DROP CONSTRAINT IF EXISTS wisp_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_channeled_actions DROP CONSTRAINT IF EXISTS wisp_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.wisp_actions DROP CONSTRAINT IF EXISTS wisp_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.vi DROP CONSTRAINT IF EXISTS vi_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_heroes DROP CONSTRAINT IF EXISTS vi_heroes_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_followers DROP CONSTRAINT IF EXISTS vi_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_events DROP CONSTRAINT IF EXISTS vi_events_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_channeled_events DROP CONSTRAINT IF EXISTS vi_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_channeled_actions DROP CONSTRAINT IF EXISTS vi_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.vi_actions DROP CONSTRAINT IF EXISTS vi_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.bold DROP CONSTRAINT IF EXISTS unique_restraint;
ALTER TABLE IF EXISTS ONLY public.turn DROP CONSTRAINT IF EXISTS turn_pkey;
ALTER TABLE IF EXISTS ONLY public.turn DROP CONSTRAINT IF EXISTS turn_name_key;
ALTER TABLE IF EXISTS ONLY public.troika DROP CONSTRAINT IF EXISTS troika_name_key;
ALTER TABLE IF EXISTS ONLY public.triggered_abilities DROP CONSTRAINT IF EXISTS triggered_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.tolerances DROP CONSTRAINT IF EXISTS tolerances_pkey;
ALTER TABLE IF EXISTS ONLY public.tinsel DROP CONSTRAINT IF EXISTS tinsel_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_heroes DROP CONSTRAINT IF EXISTS tinsel_heroes_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_events DROP CONSTRAINT IF EXISTS tinsel_events_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_continuous_events DROP CONSTRAINT IF EXISTS tinsel_continuous_events_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_continuous_channeled_events DROP CONSTRAINT IF EXISTS tinsel_continuous_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_channeled_events DROP CONSTRAINT IF EXISTS tinsel_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_channeled_actions DROP CONSTRAINT IF EXISTS tinsel_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.tinsel_actions DROP CONSTRAINT IF EXISTS tinsel_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.tendril DROP CONSTRAINT IF EXISTS tendril_name_key;
ALTER TABLE IF EXISTS ONLY public.tendril_followers DROP CONSTRAINT IF EXISTS tendril_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.tendril_events DROP CONSTRAINT IF EXISTS tendril_events_name_key;
ALTER TABLE IF EXISTS ONLY public.tendril_actions DROP CONSTRAINT IF EXISTS tendril_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.slang DROP CONSTRAINT IF EXISTS slang_pkey;
ALTER TABLE IF EXISTS ONLY public.scuttler DROP CONSTRAINT IF EXISTS scuttler_name_key;
ALTER TABLE IF EXISTS ONLY public.scuttler_followers DROP CONSTRAINT IF EXISTS scuttler_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.scinter DROP CONSTRAINT IF EXISTS scinter_name_key;
ALTER TABLE IF EXISTS ONLY public.scinter_followers DROP CONSTRAINT IF EXISTS scinter_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.scinter_events DROP CONSTRAINT IF EXISTS scinter_events_name_key;
ALTER TABLE IF EXISTS ONLY public.scinter_channeled_actions DROP CONSTRAINT IF EXISTS scinter_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.scinter_actions DROP CONSTRAINT IF EXISTS scinter_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat DROP CONSTRAINT IF EXISTS ravat_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_followers DROP CONSTRAINT IF EXISTS ravat_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_events DROP CONSTRAINT IF EXISTS ravat_events_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_continuous_events DROP CONSTRAINT IF EXISTS ravat_continuous_events_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_continuous_channeled_events DROP CONSTRAINT IF EXISTS ravat_continuous_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_channeled_events DROP CONSTRAINT IF EXISTS ravat_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_channeled_actions DROP CONSTRAINT IF EXISTS ravat_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.ravat_actions DROP CONSTRAINT IF EXISTS ravat_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.partners DROP CONSTRAINT IF EXISTS partners_name_key;
ALTER TABLE IF EXISTS ONLY public.nightmares DROP CONSTRAINT IF EXISTS nightmares_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith DROP CONSTRAINT IF EXISTS lilith_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith_followers DROP CONSTRAINT IF EXISTS lilith_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith_events DROP CONSTRAINT IF EXISTS lilith_events_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith_continuous_events DROP CONSTRAINT IF EXISTS lilith_continuous_events_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith_channeled_events DROP CONSTRAINT IF EXISTS lilith_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.lilith_actions DROP CONSTRAINT IF EXISTS lilith_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.leaders DROP CONSTRAINT IF EXISTS leaders_name_key;
ALTER TABLE IF EXISTS ONLY public.igrath DROP CONSTRAINT IF EXISTS igrath_name_key;
ALTER TABLE IF EXISTS ONLY public.igrath_followers DROP CONSTRAINT IF EXISTS igrath_followers_name_key;
ALTER TABLE IF EXISTS ONLY public.igrath_events DROP CONSTRAINT IF EXISTS igrath_events_name_key;
ALTER TABLE IF EXISTS ONLY public.igrath_continuous_events DROP CONSTRAINT IF EXISTS igrath_continuous_events_name_key;
ALTER TABLE IF EXISTS ONLY public.followers DROP CONSTRAINT IF EXISTS followers_name_key;
ALTER TABLE IF EXISTS ONLY public.events DROP CONSTRAINT IF EXISTS events_name_key;
ALTER TABLE IF EXISTS ONLY public.deck_characters DROP CONSTRAINT IF EXISTS deck_characters_name_key;
ALTER TABLE IF EXISTS ONLY public.deck_cards DROP CONSTRAINT IF EXISTS deck_cards_name_key;
ALTER TABLE IF EXISTS ONLY public.continuous DROP CONSTRAINT IF EXISTS continuous_name_key;
ALTER TABLE IF EXISTS ONLY public.constant_character_abilities DROP CONSTRAINT IF EXISTS constant_character_abilities_root_key;
ALTER TABLE IF EXISTS ONLY public.constant_character_abilities DROP CONSTRAINT IF EXISTS constant_character_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.constant_card_abilities DROP CONSTRAINT IF EXISTS constant_card_abilities_root_key;
ALTER TABLE IF EXISTS ONLY public.constant_card_abilities DROP CONSTRAINT IF EXISTS constant_card_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.constant_abilities DROP CONSTRAINT IF EXISTS constant_abilities_root_key;
ALTER TABLE IF EXISTS ONLY public.constant_abilities DROP CONSTRAINT IF EXISTS constant_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.conditional_lane_abilities DROP CONSTRAINT IF EXISTS conditional_lane_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.conditional_card_abilities DROP CONSTRAINT IF EXISTS conditional_card_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.conditional_abilities DROP CONSTRAINT IF EXISTS conditional_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.characters DROP CONSTRAINT IF EXISTS characters_name_key2;
ALTER TABLE IF EXISTS ONLY public.characters DROP CONSTRAINT IF EXISTS characters_name_key1;
ALTER TABLE IF EXISTS ONLY public.characters DROP CONSTRAINT IF EXISTS characters_name_key;
ALTER TABLE IF EXISTS ONLY public.channeled DROP CONSTRAINT IF EXISTS channeled_name_key;
ALTER TABLE IF EXISTS ONLY public.cards DROP CONSTRAINT IF EXISTS cards_pkey;
ALTER TABLE IF EXISTS ONLY public.bast DROP CONSTRAINT IF EXISTS bast_name_key;
ALTER TABLE IF EXISTS ONLY public.bast_events DROP CONSTRAINT IF EXISTS bast_events_name_key;
ALTER TABLE IF EXISTS ONLY public.bast_continuous_events DROP CONSTRAINT IF EXISTS bast_continuous_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.bast_channeled_events DROP CONSTRAINT IF EXISTS bast_channeled_events_name_key;
ALTER TABLE IF EXISTS ONLY public.bast_channeled_actions DROP CONSTRAINT IF EXISTS bast_channeled_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.bast_actions DROP CONSTRAINT IF EXISTS bast_actions_name_key;
ALTER TABLE IF EXISTS ONLY public.activities DROP CONSTRAINT IF EXISTS activities_name_key;
ALTER TABLE IF EXISTS ONLY public.activated_abilities DROP CONSTRAINT IF EXISTS activated_abilities_name_key;
ALTER TABLE IF EXISTS ONLY public.actions DROP CONSTRAINT IF EXISTS actions_name_key;
ALTER TABLE IF EXISTS public.wisp_heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_heroes ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_heroes ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_heroes ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_heroes ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.wisp ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_heroes ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_heroes ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.vi ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.turn ALTER COLUMN "order" DROP DEFAULT;
ALTER TABLE IF EXISTS public.troika ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.troika ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.troika ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.triggered_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_heroes ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_heroes ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_continuous_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_continuous_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tinsel ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.tendril ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scuttler ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.scinter ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_continuous_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_continuous_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.ravat ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.partners ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.partners ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.partners ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.nightmares ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.nightmares ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.nightmares ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.lilith ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.leaders ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.leaders ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.leaders ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.igrath ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.followers ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.followers ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.followers ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.followers ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.followers ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_heroes ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_heroes ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_heroes ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_heroes ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_heroes ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_characters ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_characters ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_characters ALTER COLUMN life DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_characters ALTER COLUMN speed DROP DEFAULT;
ALTER TABLE IF EXISTS public.deck_characters ALTER COLUMN damage DROP DEFAULT;
ALTER TABLE IF EXISTS public.continuous ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.continuous ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.constant_character_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.constant_card_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.constant_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.conditional_lane_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.conditional_card_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.conditional_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.channeled ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.channeled ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_continuous_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_continuous_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_channeled_events ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_channeled_events ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_channeled_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_channeled_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast_actions ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.bast ALTER COLUMN resolve_cost DROP DEFAULT;
ALTER TABLE IF EXISTS public.activities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.activated_abilities ALTER COLUMN complexity DROP DEFAULT;
ALTER TABLE IF EXISTS public.actions ALTER COLUMN copies DROP DEFAULT;
ALTER TABLE IF EXISTS public.actions ALTER COLUMN resolve_cost DROP DEFAULT;
DROP TABLE IF EXISTS public.wisp_continuous_events;
DROP TABLE IF EXISTS public.wisp_channeled_heroes;
DROP TABLE IF EXISTS public.wisp_heroes;
DROP TABLE IF EXISTS public.wisp_channeled_events;
DROP TABLE IF EXISTS public.wisp_events;
DROP TABLE IF EXISTS public.wisp_channeled_actions;
DROP TABLE IF EXISTS public.wisp_actions;
DROP TABLE IF EXISTS public.wisp;
DROP TABLE IF EXISTS public.vi_heroes;
DROP TABLE IF EXISTS public.vi_followers;
DROP TABLE IF EXISTS public.vi_channeled_events;
DROP TABLE IF EXISTS public.vi_events;
DROP TABLE IF EXISTS public.vi_channeled_actions;
DROP TABLE IF EXISTS public.vi_actions;
DROP TABLE IF EXISTS public.vi;
DROP SEQUENCE IF EXISTS public.turn_order_seq;
DROP TABLE IF EXISTS public.troika;
DROP TABLE IF EXISTS public.triggered_abilities;
DROP TABLE IF EXISTS public.tolerances;
DROP TABLE IF EXISTS public.tinsel_heroes;
DROP TABLE IF EXISTS public.tinsel_continuous_events;
DROP TABLE IF EXISTS public.tinsel_continuous_channeled_events;
DROP TABLE IF EXISTS public.tinsel_channeled_events;
DROP TABLE IF EXISTS public.tinsel_events;
DROP TABLE IF EXISTS public.tinsel_channeled_actions;
DROP TABLE IF EXISTS public.tinsel_actions;
DROP TABLE IF EXISTS public.tinsel;
DROP TABLE IF EXISTS public.tendril_followers;
DROP TABLE IF EXISTS public.tendril_events;
DROP TABLE IF EXISTS public.tendril_actions;
DROP TABLE IF EXISTS public.tendril;
DROP TABLE IF EXISTS public.slang;
DROP TABLE IF EXISTS public.scuttler_followers;
DROP TABLE IF EXISTS public.scuttler;
DROP TABLE IF EXISTS public.scinter_followers;
DROP TABLE IF EXISTS public.scinter_events;
DROP TABLE IF EXISTS public.scinter_channeled_actions;
DROP TABLE IF EXISTS public.scinter_actions;
DROP TABLE IF EXISTS public.scinter;
DROP TABLE IF EXISTS public.ravat_followers;
DROP TABLE IF EXISTS public.ravat_continuous_events;
DROP TABLE IF EXISTS public.ravat_continuous_channeled_events;
DROP TABLE IF EXISTS public.ravat_channeled_events;
DROP TABLE IF EXISTS public.ravat_events;
DROP TABLE IF EXISTS public.ravat_channeled_actions;
DROP TABLE IF EXISTS public.ravat_actions;
DROP TABLE IF EXISTS public.ravat;
DROP TABLE IF EXISTS public.partners;
DROP TABLE IF EXISTS public.nightmares;
DROP TABLE IF EXISTS public.lilith_followers;
DROP TABLE IF EXISTS public.lilith_continuous_events;
DROP TABLE IF EXISTS public.lilith_channeled_events;
DROP TABLE IF EXISTS public.lilith_events;
DROP TABLE IF EXISTS public.lilith_actions;
DROP TABLE IF EXISTS public.lilith;
DROP TABLE IF EXISTS public.igrath_followers;
DROP TABLE IF EXISTS public.igrath_continuous_events;
DROP TABLE IF EXISTS public.igrath_events;
DROP TABLE IF EXISTS public.igrath;
DROP VIEW IF EXISTS public.glossary;
DROP TABLE IF EXISTS public.turn;
DROP TABLE IF EXISTS public.followers;
DROP TABLE IF EXISTS public.deck_heroes;
DROP TABLE IF EXISTS public.deck_characters;
DROP TABLE IF EXISTS public.constant_character_abilities;
DROP TABLE IF EXISTS public.constant_card_abilities;
DROP TABLE IF EXISTS public.constant_abilities;
DROP TABLE IF EXISTS public.conditional_lane_abilities;
DROP TABLE IF EXISTS public.conditional_card_abilities;
DROP TABLE IF EXISTS public.conditional_abilities;
DROP VIEW IF EXISTS public.completed;
DROP VIEW IF EXISTS public.card_types;
DROP VIEW IF EXISTS public.bold_regexp;
DROP TABLE IF EXISTS public.zones;
DROP TABLE IF EXISTS public.leaders;
DROP TABLE IF EXISTS public.heroes;
DROP TABLE IF EXISTS public.characters;
DROP VIEW IF EXISTS public.card_traits;
DROP VIEW IF EXISTS public.inheritence;
DROP TABLE IF EXISTS public.bold;
DROP TABLE IF EXISTS public.bast_continuous_events;
DROP TABLE IF EXISTS public.continuous;
DROP TABLE IF EXISTS public.bast_channeled_events;
DROP TABLE IF EXISTS public.bast_events;
DROP TABLE IF EXISTS public.bast_channeled_actions;
DROP TABLE IF EXISTS public.channeled;
DROP TABLE IF EXISTS public.bast_actions;
DROP TABLE IF EXISTS public.bast;
DROP TABLE IF EXISTS public.activities;
DROP TABLE IF EXISTS public.activated_abilities;
DROP TABLE IF EXISTS public.actions;
DROP TABLE IF EXISTS public.events;
DROP VIEW IF EXISTS public.ability_types;
DROP VIEW IF EXISTS public.abilities_with_type;
DROP FUNCTION IF EXISTS public.type(card cards);
DROP FUNCTION IF EXISTS public.type(ability abilities);
DROP FUNCTION IF EXISTS public.text(ability abilities);
DROP FUNCTION IF EXISTS public.test2(src text, reg text);
DROP FUNCTION IF EXISTS public.test(card cards);
DROP FUNCTION IF EXISTS public.supertypes(card cards);
DROP FUNCTION IF EXISTS public.subtype(ability abilities);
DROP FUNCTION IF EXISTS public.speed_b(card cards);
DROP FUNCTION IF EXISTS public.speed(card cards);
DROP FUNCTION IF EXISTS public.short_b(card cards);
DROP FUNCTION IF EXISTS public.short(card cards);
DROP FUNCTION IF EXISTS public.rules(ability abilities);
DROP FUNCTION IF EXISTS public.resolve_b(card cards);
DROP FUNCTION IF EXISTS public.regexp(card cards);
DROP FUNCTION IF EXISTS public.rarity(card cards);
DROP FUNCTION IF EXISTS public.long_b(card cards);
DROP FUNCTION IF EXISTS public.long(card cards);
DROP FUNCTION IF EXISTS public.life_b(card cards);
DROP FUNCTION IF EXISTS public.life(card cards);
DROP FUNCTION IF EXISTS public.leader(card deck_cards);
DROP TABLE IF EXISTS public.deck_cards;
DROP FUNCTION IF EXISTS public.leader(card cards);
DROP FUNCTION IF EXISTS public.getcard2("table" text, name text);
DROP FUNCTION IF EXISTS public.get_all_children(tblname text);
DROP FUNCTION IF EXISTS public.flavor_b(card cards);
DROP FUNCTION IF EXISTS public.faction(card cards);
DROP FUNCTION IF EXISTS public.damage_b(card cards);
DROP FUNCTION IF EXISTS public.damage(card cards);
DROP FUNCTION IF EXISTS public.cost(card cards);
DROP TABLE IF EXISTS public.cards;
DROP FUNCTION IF EXISTS public.cost(ability abilities);
DROP FUNCTION IF EXISTS public.constant_character_ability_rules();
DROP FUNCTION IF EXISTS public.conditional_lane_ability_rules();
DROP FUNCTION IF EXISTS public.conditional_card_ability_rules();
DROP FUNCTION IF EXISTS public.condition(ability abilities);
DROP TABLE IF EXISTS public.abilities;
DROP FUNCTION IF EXISTS public.caseless(t text);
DROP FUNCTION IF EXISTS public.card_ability_fk();
DROP FUNCTION IF EXISTS public.break(t text);
DROP FUNCTION IF EXISTS public.bast_insert_action();
DROP FUNCTION IF EXISTS public.activity_rules();
DROP SCHEMA IF EXISTS public;
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET search_path = public, pg_catalog;

--
-- Name: activity_rules(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION activity_rules() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        IF NEW.cost IS NULL THEN
            NEW.rules = (SELECT 'To ' || NEW.root);
        ELSE
            NEW.rules = (SELECT 'To pay ' || NEW.cost || ' to ' || NEW.root);
        END IF;
        RETURN NEW;
    END; $$;


ALTER FUNCTION public.activity_rules() OWNER TO postgres;

--
-- Name: bast_insert_action(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION bast_insert_action() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.type = 'Action' THEN
        INSERT INTO bast_actions
        (name, resolve, short, reminder, flavor, resolve_cost, copies) VALUES (
            NEW.name,
            NEW.resolve,
            NEW.short,
            NEW.reminder,
            NEW.flavor,
            NEW.resolve_cost,
            NEW.copies
        );
    END IF;
    RETURN NULL;
END; $$;


ALTER FUNCTION public.bast_insert_action() OWNER TO postgres;

--
-- Name: break(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION break(t text) RETURNS text[]
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN string_to_array(t, E'\n');
END; $$;


ALTER FUNCTION public.break(t text) OWNER TO postgres;

--
-- Name: card_ability_fk(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION card_ability_fk() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
	If NEW.card = (SELECT name from cards where name=New.card) AND
	NEW.ability = (SELECT name from abilities where name=NEW.ability)
	THEN
		RETURN NEW;
	ELSE 
		RETURN NULL;
	END IF;
END; $$;


ALTER FUNCTION public.card_ability_fk() OWNER TO postgres;

--
-- Name: caseless(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION caseless(t text) RETURNS text
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN '[' || lower(substring(t from 1 for 1))
		|| upper(substring(t from 1 for 1)) || ']'
		|| substring(t from 2);
END; $$;


ALTER FUNCTION public.caseless(t text) OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE abilities (
    name text NOT NULL,
    root text,
    complexity integer DEFAULT 1 NOT NULL
);


ALTER TABLE abilities OWNER TO postgres;

--
-- Name: TABLE abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE abilities IS 'A line of rules text. Abilities appear in card''s short text and the rulebook.';


--
-- Name: condition(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION condition(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN (SELECT condition from conditional_abilities WHERE name=ability.name)::TEXT;
END; $$;


ALTER FUNCTION public.condition(ability abilities) OWNER TO postgres;

--
-- Name: conditional_card_ability_rules(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION conditional_card_ability_rules() RETURNS trigger
    LANGUAGE plpgsql
    AS $$    BEGIN
        NEW.rules = (SELECT 'Cards are ' || lower(NEW.name) || ' if ' || NEW.condition);
        RETURN NEW;
    END; $$;


ALTER FUNCTION public.conditional_card_ability_rules() OWNER TO postgres;

--
-- Name: conditional_lane_ability_rules(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION conditional_lane_ability_rules() RETURNS trigger
    LANGUAGE plpgsql
    AS $$    BEGIN
        IF NEW.condition ~ 'they' THEN
            NEW.rules = (SELECT 'Lanes are ' || lower(NEW.name) || ' if ' || NEW.condition);
        ELSE
            NEW.rules = (SELECT 'A lane is ' || lower(NEW.name) || ' if ' || NEW.condition);
        END IF;
        RETURN NEW;
    END; $$;


ALTER FUNCTION public.conditional_lane_ability_rules() OWNER TO postgres;

--
-- Name: constant_character_ability_rules(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION constant_character_ability_rules() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        IF NEW.root ~ '^''t' THEN
            NEW.rules = (SELECT 'Characters with ' || NEW.name || ' can' || NEW.root);
        ELSE
            NEW.rules = (SELECT 'Characters with ' || NEW.name || ' can ' || NEW.root);
        END IF;
        RETURN NEW;
    END; $$;


ALTER FUNCTION public.constant_character_ability_rules() OWNER TO postgres;

--
-- Name: cost(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION cost(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	cost TEXT := (SELECT cost from activated_abilities WHERE name=ability.name);
BEGIN
	IF cost IS NULL THEN
		IF EXISTS(SELECT name from activated_abilities WHERE name=ability.name) THEN
			RETURN '';
		ELSE
			RETURN NULL;
		END IF;
	END IF;
	RETURN cost;
END; $$;


ALTER FUNCTION public.cost(ability abilities) OWNER TO postgres;

--
-- Name: cards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE cards (
    name text NOT NULL,
    resolve character(2),
    abilities text[],
    reminder text,
    flavor text
);


ALTER TABLE cards OWNER TO postgres;

--
-- Name: COLUMN cards.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN cards.name IS 'A unique identifier';


--
-- Name: COLUMN cards.resolve; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN cards.resolve IS 'A resource generated by cards that can be spent to play other cards.';


--
-- Name: cost(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION cost(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT resolve_cost from deck_cards where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.cost(card cards) OWNER TO postgres;

--
-- Name: damage(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION damage(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT damage from characters where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.damage(card cards) OWNER TO postgres;

--
-- Name: damage_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION damage_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT damage_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.damage_b(card cards) OWNER TO postgres;

--
-- Name: faction(cards); Type: FUNCTION; Schema: public; Owner: sbrow
--

CREATE FUNCTION faction(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
BEGIN
	IF type(card)='Leader' THEN
		RETURN INITCAP((SELECT tableoid::regclass::TEXT from public.cards 
    WHERE "name" = card.name LIMIT 1)::TEXT);
	ELSE
		 RETURN NULL;
	END IF;
END; $$;


ALTER FUNCTION public.faction(card cards) OWNER TO sbrow;

--
-- Name: flavor_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION flavor_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT flavor_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.flavor_b(card cards) OWNER TO postgres;

--
-- Name: get_all_children(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION get_all_children(tblname text) RETURNS TABLE(parent text, child text, description text)
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN QUERY (WITH RECURSIVE f(parent, child) AS (
		SELECT i.parent, i.child FROM inheritence i WHERE i.parent=tblname
		UNION ALL
		SELECT p.parent, p.child FROM f pr, inheritence p
			WHERE p.parent=pr.child
			AND (SELECT count(*) from leaders where pr.child ~ lower(name) OR p.child ~ lower(name))=0)
	SELECT DISTINCT p.relowner::TEXT, f.child::TEXT, obj_description(p.oid) as description FROM f, pg_class p WHERE p.relname=f.child
	UNION
	SELECT relowner::TEXT, relname::TEXT, obj_description(oid) as description from pg_class where relname=tblname
	ORDER BY child ASC);
END; $$;


ALTER FUNCTION public.get_all_children(tblname text) OWNER TO postgres;

--
-- Name: getcard2(text, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION getcard2("table" text, name text) RETURNS record
    LANGUAGE plpgsql
    AS $$
BEGIN
	EXECUTE 'SELECT * FROM public.' || quote_ident(public.getCard("name")) || ' WHERE "name"=' || quote_nullable("name") || ';';
END; $$;


ALTER FUNCTION public.getcard2("table" text, name text) OWNER TO postgres;

--
-- Name: leader(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION leader(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT deck_cards.leader from deck_cards where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.leader(card cards) OWNER TO postgres;

--
-- Name: deck_cards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE deck_cards (
    resolve_cost text DEFAULT '1'::text NOT NULL,
    copies integer DEFAULT 3 NOT NULL,
    CONSTRAINT deck_cards_copies_check CHECK (((copies > 0) AND (copies <= 3)))
)
INHERITS (cards);


ALTER TABLE deck_cards OWNER TO postgres;

--
-- Name: TABLE deck_cards; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE deck_cards IS 'Cards that appear in a leader''s deck.';


--
-- Name: COLUMN deck_cards.resolve_cost; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN deck_cards.resolve_cost IS 'The amount of resolve that must be spent to play a card from hand. Can be paid with resolve from one or more sources.';


--
-- Name: leader(deck_cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION leader(card deck_cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	tbl TEXT := (select tableoid::regclass::text as tbl from deck_cards where name=card.name LIMIT 1)::TEXT;
	arr TEXT[] := regexp_split_to_array(tbl, '_');
BEGIN
	RETURN INITCAP(arr[1]);
END; $$;


ALTER FUNCTION public.leader(card deck_cards) OWNER TO postgres;

--
-- Name: life(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION life(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT life from characters where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.life(card cards) OWNER TO postgres;

--
-- Name: life_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION life_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT life_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.life_b(card cards) OWNER TO postgres;

--
-- Name: long(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION long(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	long TEXT := '';
BEGIN
	IF card.reminder IS NOT NULL THEN
		long := long || card.reminder || E'\r';
	END IF;
	long := long || array_to_string(array(
		SELECT rules FROM glossary
		WHERE (supertypes(card) || ' Card' ~* name 
			OR short(card) ~* name)
		AND complexity > 0
		ORDER by complexity DESC)	
	, E'\r');
	RETURN long;
END; $$;


ALTER FUNCTION public.long(card cards) OWNER TO postgres;

--
-- Name: long_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION long_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT reminder_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.long_b(card cards) OWNER TO postgres;

--
-- Name: rarity(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION rarity(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT copies from deck_cards where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.rarity(card cards) OWNER TO postgres;

--
-- Name: regexp(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION regexp(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	r text[] := string_to_array(regexp_replace(regexp_replace(regexp_replace(short(card),
		'\+', '\\+', 'gi'),
		'[\r\n]', ' '),
		'\{', '\\{'), ' ');
BEGIN
	RETURN '(' ||array_to_string(array(SELECT DISTINCT unnest from unnest(r) WHERE
		EXISTS(select FROM bold_regexp WHERE unnest ~* name)), ')|(') || ')';
END; $$;


ALTER FUNCTION public.regexp(card cards) OWNER TO postgres;

--
-- Name: resolve_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION resolve_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT resolve_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.resolve_b(card cards) OWNER TO postgres;

--
-- Name: rules(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION rules(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	tbl TEXT := (SELECT tableoid::regclass::TEXT FROM abilities where name=ability.name LIMIT 1);
	rules_text TEXT := '';
BEGIN
	rules_text := (
		SELECT CASE
		WHEN tbl = 'abilities' THEN ability.root
		WHEN cost(ability) = '' THEN ability. name || ' means "to ' || ability.root || '"'
		WHEN cost(ability) IS NOT NULL AND NOT cost(ability) = '' THEN 'To pay ' || cost(ability) || ' to ' || ability.root
		WHEN condition(ability) IS NOT NULL THEN text(ability)
		WHEN type(ability) = 'Constant' THEN subtype(ability) || 's with ' || ability.name || ' can' || ability.root
		END
	);
	RETURN rules_text;
END; $$;


ALTER FUNCTION public.rules(ability abilities) OWNER TO postgres;

--
-- Name: short(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION short(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
BEGIN
	RETURN array_to_string(card.abilities, E'\n');
END; $$;


ALTER FUNCTION public.short(card cards) OWNER TO postgres;

--
-- Name: short_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION short_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT short_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.short_b(card cards) OWNER TO postgres;

--
-- Name: speed(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION speed(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT speed from characters where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.speed(card cards) OWNER TO postgres;

--
-- Name: speed_b(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION speed_b(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
BEGIN
	RETURN (SELECT speed_b from leaders where name = card.name)::TEXT;
END; $$;


ALTER FUNCTION public.speed_b(card cards) OWNER TO postgres;

--
-- Name: subtype(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION subtype(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	arr text[] := string_to_array((SELECT tableoid::regclass::TEXT from abilities where name=ability.name LIMIT 1), '_');
BEGIN
	RETURN INITCAP(arr[array_length(arr, 1)-1]);
END; $$;


ALTER FUNCTION public.subtype(ability abilities) OWNER TO postgres;

--
-- Name: supertypes(cards); Type: FUNCTION; Schema: public; Owner: sbrow
--

CREATE FUNCTION supertypes(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $_$ 
  DECLARE 
    arr text[]; 
    table_name text; 
    len int; 
 BEGIN 
   table_name := (SELECT tableoid::regclass::TEXT from public.cards 
    WHERE "name" = card.name LIMIT 1); 
  arr := regexp_split_to_array(table_name, '_'); 
  len := array_length(arr, 1); 
  CASE 
    WHEN len <= 2 THEN
      table_name := regexp_replace(table_name, '^[^_]*_', '');
      CASE
        WHEN table_name = 'heroes' THEN
          RETURN 'Deck';
        WHEN table_name = 'guests' THEN
          RETURN 'Partner';
        ELSE
          RETURN 'Leader';
      END CASE;
    ELSE 
      table_name := regexp_replace(table_name, '^[^_]*_', ''); 
      table_name := regexp_replace(table_name, '_[^_]*$', ''); 
      table_name := INITCAP(trim(regexp_replace(table_name, '_', ',')));
      RETURN table_name; 
  END CASE; 
 END; $_$;


ALTER FUNCTION public.supertypes(card cards) OWNER TO sbrow;

--
-- Name: test(cards); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION test(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	abilities TEXT ARRAY := array(SELECT abilities.rules
	FROM abilities
	WHERE array_to_string(card.short,',') ~ abilities.name
	ORDER BY abilities.name ASC);
BEGIN
	RETURN array_to_string(abilities, E'\n');
END; $$;


ALTER FUNCTION public.test(card cards) OWNER TO postgres;

--
-- Name: test2(text, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION test2(src text, reg text) RETURNS text
    LANGUAGE plpgsql
    AS $$
BEGIN
	return array_to_string(array(select distinct regexp_matches(src, reg, 'g')), '|');
END; $$;


ALTER FUNCTION public.test2(src text, reg text) OWNER TO postgres;

--
-- Name: text(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION text(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $_$
DECLARE
	str TEXT := ability.root;
BEGIN
	str := regexp_replace(str, '\${name}', lower(ability.name));
	str := regexp_replace(str, '\${type}', subtype(ability));
	str := regexp_replace(str, '\${condition}', condition(ability));
	RETURN str;
END; $_$;


ALTER FUNCTION public.text(ability abilities) OWNER TO postgres;

--
-- Name: type(abilities); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION type(ability abilities) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
	arr text[] := string_to_array((SELECT tableoid::regclass::TEXT from abilities where name=ability.name LIMIT 1), '_');
BEGIN
	RETURN INITCAP(arr[1]);
END; $$;


ALTER FUNCTION public.type(ability abilities) OWNER TO postgres;

--
-- Name: type(cards); Type: FUNCTION; Schema: public; Owner: sbrow
--

CREATE FUNCTION type(card cards) RETURNS text
    LANGUAGE plpgsql
    AS $$ 
 DECLARE 
  arr text[]; 
  n text; 
  len int; 
 BEGIN 
   n := (SELECT tableoid::regclass::TEXT from public.cards 
    WHERE "name" = card.name LIMIT 1); 
   arr := regexp_split_to_array(n, '_'); 
  len := array_length(arr, 1); 
  IF arr[len] = 'troika' OR arr[len] = 'nightmares'  or arr[len] = 'heroes' THEN
    return 'Hero';
  END IF;
  return INITCAP(trim(trailing 'es' from arr[len])); 
 END; $$;


ALTER FUNCTION public.type(card cards) OWNER TO sbrow;

--
-- Name: abilities_with_type; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW abilities_with_type AS
 SELECT ((abilities.tableoid)::regclass)::text AS parent,
    abilities.name,
    abilities.root AS rules
   FROM abilities
  ORDER BY abilities.name, ((abilities.tableoid)::regclass)::text;


ALTER TABLE abilities_with_type OWNER TO postgres;

--
-- Name: ability_types; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW ability_types AS
 SELECT get_all_children.parent,
    get_all_children.child,
    get_all_children.description
   FROM get_all_children('abilities'::text) get_all_children(parent, child, description);


ALTER TABLE ability_types OWNER TO postgres;

--
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE events (
)
INHERITS (deck_cards);


ALTER TABLE events OWNER TO postgres;

--
-- Name: TABLE events; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE events IS 'Perform an effect, and are then discarded.';


--
-- Name: actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE actions (
)
INHERITS (events);


ALTER TABLE actions OWNER TO postgres;

--
-- Name: TABLE actions; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE actions IS 'May be played at any time for a temporary effect.';


--
-- Name: activated_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE activated_abilities (
    cost text
)
INHERITS (abilities);


ALTER TABLE activated_abilities OWNER TO postgres;

--
-- Name: TABLE activated_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE activated_abilities IS 'Performs a temporary effect whenever its activation cost is paid. Activated abilities appear as "Cost: Effect".';


--
-- Name: COLUMN activated_abilities.cost; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN activated_abilities.cost IS 'The price that must be paid in order to activate this ability.';


--
-- Name: activities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE activities (
)
INHERITS (activated_abilities);


ALTER TABLE activities OWNER TO postgres;

--
-- Name: bast; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast (
)
INHERITS (deck_cards);


ALTER TABLE bast OWNER TO postgres;

--
-- Name: bast_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast_actions (
)
INHERITS (bast, actions);


ALTER TABLE bast_actions OWNER TO postgres;

--
-- Name: channeled; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE channeled (
)
INHERITS (deck_cards);


ALTER TABLE channeled OWNER TO postgres;

--
-- Name: TABLE channeled; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE channeled IS 'must be played with their leader''s resolve.';


--
-- Name: bast_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast_channeled_actions (
)
INHERITS (bast_actions, channeled);


ALTER TABLE bast_channeled_actions OWNER TO postgres;

--
-- Name: bast_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast_events (
)
INHERITS (bast, events);


ALTER TABLE bast_events OWNER TO postgres;

--
-- Name: bast_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast_channeled_events (
)
INHERITS (bast_events, channeled);


ALTER TABLE bast_channeled_events OWNER TO postgres;

--
-- Name: continuous; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE continuous (
)
INHERITS (deck_cards);


ALTER TABLE continuous OWNER TO postgres;

--
-- Name: TABLE continuous; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE continuous IS 'Remain in play until the end of the game, or removed by another effect.';


--
-- Name: bast_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bast_continuous_events (
)
INHERITS (continuous, bast_events);


ALTER TABLE bast_continuous_events OWNER TO postgres;

--
-- Name: bold; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE bold (
    regex text NOT NULL
);


ALTER TABLE bold OWNER TO postgres;

--
-- Name: TABLE bold; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE bold IS 'Miscellaneous things that should appear bold on cards';


--
-- Name: inheritence; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW inheritence AS
 SELECT p.relname AS parent,
    c.relname AS child
   FROM ((pg_inherits
     JOIN pg_class c ON ((pg_inherits.inhrelid = c.oid)))
     JOIN pg_class p ON ((pg_inherits.inhparent = p.oid)));


ALTER TABLE inheritence OWNER TO postgres;

--
-- Name: card_traits; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW card_traits WITH (security_barrier='false') AS
 SELECT c.table_name AS tble,
    c.column_name AS name,
    pgd.description AS rules
   FROM ((pg_statio_all_tables st
     JOIN pg_description pgd ON ((pgd.objoid = st.relid)))
     JOIN information_schema.columns c ON (((pgd.objsubid = (c.ordinal_position)::integer) AND ((c.table_schema)::text = 'public'::text) AND ((c.table_name)::name = st.relname))))
  WHERE (EXISTS ( SELECT inheritence.parent,
            inheritence.child
           FROM inheritence));


ALTER TABLE card_traits OWNER TO postgres;

--
-- Name: characters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE characters (
    damage integer DEFAULT 1 NOT NULL,
    speed integer DEFAULT 1 NOT NULL,
    life integer DEFAULT 1 NOT NULL
)
INHERITS (cards);


ALTER TABLE characters OWNER TO postgres;

--
-- Name: TABLE characters; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE characters IS 'Cards that be engaged in skirmishes';


--
-- Name: COLUMN characters.damage; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN characters.damage IS 'The amount of life a character removes when skirmishing.';


--
-- Name: COLUMN characters.speed; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN characters.speed IS 'Speed can be spent to attack, intercept, or redeploy.';


--
-- Name: COLUMN characters.life; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN characters.life IS 'The amount of damage a character can survive.';


--
-- Name: heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE heroes (
)
INHERITS (characters);


ALTER TABLE heroes OWNER TO postgres;

--
-- Name: TABLE heroes; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE heroes IS 'Characters with high life totals that lose life for damage they''re dealt.';


--
-- Name: leaders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE leaders (
    banner character(6) NOT NULL,
    indicator character(6) NOT NULL,
    resolve_b character(2),
    speed_b integer,
    damage_b integer,
    life_b character(2),
    short_b text,
    reminder_b text,
    flavor_b text
)
INHERITS (heroes);


ALTER TABLE leaders OWNER TO postgres;

--
-- Name: TABLE leaders; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE leaders IS 'A hero with a 20 card deck.';


--
-- Name: zones; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE zones (
    name character varying(15) NOT NULL,
    definition text
);


ALTER TABLE zones OWNER TO postgres;

--
-- Name: TABLE zones; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE zones IS 'A discrete area where cards can be placed during a game.';


--
-- Name: bold_regexp; Type: VIEW; Schema: public; Owner: sbrow
--

CREATE VIEW bold_regexp AS
 SELECT abilities.name
   FROM abilities
  WHERE (NOT (abilities.name ~* 'gain|lose'::text))
UNION
 SELECT bold.regex AS name
   FROM bold
UNION
 SELECT card_traits.name
   FROM card_traits
UNION
 SELECT rtrim((inheritence.child)::text, 'es'::text) AS name
   FROM inheritence
  WHERE (NOT (inheritence.parent = 'cards'::name))
UNION
 SELECT faction((leaders.*)::cards) AS name
   FROM leaders
UNION
 SELECT leaders.name
   FROM leaders
UNION
 SELECT zones.name
   FROM zones;


ALTER TABLE bold_regexp OWNER TO sbrow;

--
-- Name: card_types; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW card_types WITH (security_barrier='false') AS
 SELECT t.parent,
    initcap(rtrim(regexp_replace(t.child, '_'::text, ' '::text), 'es'::text)) AS child,
    t.description
   FROM ( SELECT get_all_children.parent,
            get_all_children.child,
            get_all_children.description
           FROM get_all_children('cards'::text) get_all_children(parent, child, description)) t;


ALTER TABLE card_types OWNER TO postgres;

--
-- Name: completed; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW completed AS
 SELECT cards.name,
    short(cards.*) AS short
   FROM cards
  WHERE (cards.name ~ (('('::text || array_to_string(((((((('{'::text || 'Chaotic Blast,Combust,Ignite,Rush of Anger,Savage Melee,Scorching Strike,Searing Focus'::text) || ',Drastic Measures'::text) || ',Diligent Research,Interfering Librarian,Tower Guard'::text) || ',Gruesome Display,Hunt,Lacerate'::text) || 'Master Manipulations,Nainso,Troika Assassin'::text) || '}'::text))::text[], ')|('::text)) || ')'::text));


ALTER TABLE completed OWNER TO postgres;

--
-- Name: conditional_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE conditional_abilities (
    condition text NOT NULL
)
INHERITS (abilities);


ALTER TABLE conditional_abilities OWNER TO postgres;

--
-- Name: TABLE conditional_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE conditional_abilities IS 'Performs a constant effect as long as its condition is met. Conditional Abilities appear as "Condtion- Effect".';


--
-- Name: conditional_card_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE conditional_card_abilities (
)
INHERITS (conditional_abilities);


ALTER TABLE conditional_card_abilities OWNER TO postgres;

--
-- Name: TABLE conditional_card_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE conditional_card_abilities IS 'Conditional abilities that only affect cards';


--
-- Name: conditional_lane_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE conditional_lane_abilities (
)
INHERITS (conditional_abilities);


ALTER TABLE conditional_lane_abilities OWNER TO postgres;

--
-- Name: TABLE conditional_lane_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE conditional_lane_abilities IS 'Conditional abilities that only affect lanes';


--
-- Name: constant_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE constant_abilities (
    rt_text text
)
INHERITS (abilities);


ALTER TABLE constant_abilities OWNER TO postgres;

--
-- Name: TABLE constant_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE constant_abilities IS 'Performing an effect at all times. i.e. "Ambush", or the "Last Stand"';


--
-- Name: constant_card_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE constant_card_abilities (
)
INHERITS (constant_abilities);


ALTER TABLE constant_card_abilities OWNER TO postgres;

--
-- Name: TABLE constant_card_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE constant_card_abilities IS 'Constant abilities that only appear on cards.';


--
-- Name: constant_character_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE constant_character_abilities (
)
INHERITS (constant_card_abilities);


ALTER TABLE constant_character_abilities OWNER TO postgres;

--
-- Name: TABLE constant_character_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE constant_character_abilities IS 'Constant abilities that only appear on characters.';


--
-- Name: deck_characters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE deck_characters (
)
INHERITS (characters, deck_cards);


ALTER TABLE deck_characters OWNER TO postgres;

--
-- Name: deck_heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE deck_heroes (
)
INHERITS (heroes, deck_characters);


ALTER TABLE deck_heroes OWNER TO postgres;

--
-- Name: followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE followers (
)
INHERITS (deck_characters);


ALTER TABLE followers OWNER TO postgres;

--
-- Name: TABLE followers; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE followers IS 'Typically have less life than heroes, but it regenerates at the end of each turn.';


--
-- Name: turn; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE turn (
    "order" integer NOT NULL,
    name text NOT NULL,
    type text NOT NULL,
    definition text,
    CONSTRAINT type CHECK (((type = 'Phase'::text) OR (type = 'Step'::text)))
);


ALTER TABLE turn OWNER TO postgres;

--
-- Name: TABLE turn; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE turn IS 'A sequence of steps and phases.';


--
-- Name: glossary; Type: VIEW; Schema: public; Owner: postgres
--

CREATE VIEW glossary AS
 SELECT t.name,
    t.rules,
    t.complexity
   FROM ( SELECT abilities.name,
            rules(abilities.*) AS rules,
            abilities.complexity
           FROM abilities
        UNION
         SELECT card_traits.name,
            card_traits.rules,
            1
           FROM card_traits
        UNION
         SELECT (card_types.child || ' Card'::text) AS name,
            ((card_types.child || ' cards '::text) || card_types.description),
            ((card_types.child = 'Channeled'::text))::integer AS int4
           FROM card_types
        UNION
         SELECT (ability_types.child || ' Ability'::text),
            ability_types.description,
            0
           FROM ability_types
        UNION
         SELECT concat(turn.name, ' ', turn.type) AS concat,
            turn.definition,
            0
           FROM turn
        UNION
         SELECT concat(zones.name, ' Zone') AS concat,
            zones.definition,
            0
           FROM zones) t
  ORDER BY t.name;


ALTER TABLE glossary OWNER TO postgres;

--
-- Name: igrath; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE igrath (
)
INHERITS (deck_cards);


ALTER TABLE igrath OWNER TO postgres;

--
-- Name: igrath_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE igrath_events (
)
INHERITS (igrath, events);


ALTER TABLE igrath_events OWNER TO postgres;

--
-- Name: igrath_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE igrath_continuous_events (
)
INHERITS (igrath_events, continuous);


ALTER TABLE igrath_continuous_events OWNER TO postgres;

--
-- Name: igrath_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE igrath_followers (
)
INHERITS (igrath, followers);


ALTER TABLE igrath_followers OWNER TO postgres;

--
-- Name: lilith; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith (
)
INHERITS (deck_cards);


ALTER TABLE lilith OWNER TO postgres;

--
-- Name: lilith_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith_actions (
)
INHERITS (lilith, actions);


ALTER TABLE lilith_actions OWNER TO postgres;

--
-- Name: lilith_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith_events (
)
INHERITS (lilith, events);


ALTER TABLE lilith_events OWNER TO postgres;

--
-- Name: lilith_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith_channeled_events (
)
INHERITS (lilith_events, channeled);


ALTER TABLE lilith_channeled_events OWNER TO postgres;

--
-- Name: lilith_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith_continuous_events (
)
INHERITS (lilith_events, continuous);


ALTER TABLE lilith_continuous_events OWNER TO postgres;

--
-- Name: lilith_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE lilith_followers (
)
INHERITS (lilith, followers);


ALTER TABLE lilith_followers OWNER TO postgres;

--
-- Name: nightmares; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE nightmares (
)
INHERITS (leaders);


ALTER TABLE nightmares OWNER TO postgres;

--
-- Name: TABLE nightmares; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE nightmares IS 'A faction.';


--
-- Name: partners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE partners (
)
INHERITS (heroes);


ALTER TABLE partners OWNER TO postgres;

--
-- Name: TABLE partners; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE partners IS 'A hero without a deck.';


--
-- Name: ravat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat (
)
INHERITS (deck_cards);


ALTER TABLE ravat OWNER TO postgres;

--
-- Name: ravat_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_actions (
)
INHERITS (ravat, actions);


ALTER TABLE ravat_actions OWNER TO postgres;

--
-- Name: ravat_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_channeled_actions (
)
INHERITS (ravat_actions, channeled);


ALTER TABLE ravat_channeled_actions OWNER TO postgres;

--
-- Name: ravat_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_events (
)
INHERITS (ravat, events);


ALTER TABLE ravat_events OWNER TO postgres;

--
-- Name: ravat_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_channeled_events (
)
INHERITS (ravat_events, channeled);


ALTER TABLE ravat_channeled_events OWNER TO postgres;

--
-- Name: ravat_continuous_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_continuous_channeled_events (
)
INHERITS (ravat_channeled_events, continuous);


ALTER TABLE ravat_continuous_channeled_events OWNER TO postgres;

--
-- Name: ravat_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_continuous_events (
)
INHERITS (ravat_events, continuous);


ALTER TABLE ravat_continuous_events OWNER TO postgres;

--
-- Name: ravat_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE ravat_followers (
)
INHERITS (ravat, followers);


ALTER TABLE ravat_followers OWNER TO postgres;

--
-- Name: scinter; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scinter (
)
INHERITS (deck_cards);


ALTER TABLE scinter OWNER TO postgres;

--
-- Name: scinter_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scinter_actions (
)
INHERITS (scinter, actions);


ALTER TABLE scinter_actions OWNER TO postgres;

--
-- Name: scinter_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scinter_channeled_actions (
)
INHERITS (scinter_actions, channeled);


ALTER TABLE scinter_channeled_actions OWNER TO postgres;

--
-- Name: scinter_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scinter_events (
)
INHERITS (scinter, events);


ALTER TABLE scinter_events OWNER TO postgres;

--
-- Name: scinter_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scinter_followers (
)
INHERITS (scinter, followers);


ALTER TABLE scinter_followers OWNER TO postgres;

--
-- Name: scuttler; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scuttler (
)
INHERITS (deck_cards);


ALTER TABLE scuttler OWNER TO postgres;

--
-- Name: scuttler_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE scuttler_followers (
)
INHERITS (scuttler, followers);


ALTER TABLE scuttler_followers OWNER TO postgres;

--
-- Name: slang; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE slang (
    name text NOT NULL,
    description text
);


ALTER TABLE slang OWNER TO postgres;

--
-- Name: tendril; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tendril (
)
INHERITS (deck_cards);


ALTER TABLE tendril OWNER TO postgres;

--
-- Name: tendril_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tendril_actions (
)
INHERITS (tendril, actions);


ALTER TABLE tendril_actions OWNER TO postgres;

--
-- Name: tendril_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tendril_events (
)
INHERITS (tendril, events);


ALTER TABLE tendril_events OWNER TO postgres;

--
-- Name: tendril_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tendril_followers (
)
INHERITS (tendril, followers);


ALTER TABLE tendril_followers OWNER TO postgres;

--
-- Name: tinsel; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel (
)
INHERITS (deck_cards);


ALTER TABLE tinsel OWNER TO postgres;

--
-- Name: tinsel_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_actions (
)
INHERITS (tinsel, actions);


ALTER TABLE tinsel_actions OWNER TO postgres;

--
-- Name: tinsel_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_channeled_actions (
)
INHERITS (tinsel_actions, channeled);


ALTER TABLE tinsel_channeled_actions OWNER TO postgres;

--
-- Name: tinsel_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_events (
)
INHERITS (tinsel, events);


ALTER TABLE tinsel_events OWNER TO postgres;

--
-- Name: tinsel_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_channeled_events (
)
INHERITS (tinsel_events, channeled);


ALTER TABLE tinsel_channeled_events OWNER TO postgres;

--
-- Name: tinsel_continuous_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_continuous_channeled_events (
)
INHERITS (tinsel_channeled_events, continuous);


ALTER TABLE tinsel_continuous_channeled_events OWNER TO postgres;

--
-- Name: tinsel_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_continuous_events (
)
INHERITS (tinsel_events, continuous);


ALTER TABLE tinsel_continuous_events OWNER TO postgres;

--
-- Name: tinsel_heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tinsel_heroes (
)
INHERITS (heroes, tinsel);


ALTER TABLE tinsel_heroes OWNER TO postgres;

--
-- Name: tolerances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tolerances (
    name text NOT NULL,
    px integer DEFAULT 0 NOT NULL
);


ALTER TABLE tolerances OWNER TO postgres;

--
-- Name: triggered_abilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE triggered_abilities (
)
INHERITS (abilities);


ALTER TABLE triggered_abilities OWNER TO postgres;

--
-- Name: TABLE triggered_abilities; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE triggered_abilities IS 'Performs an effect whenever its condition is met. Triggered abilities appear as "Condition- Effect".';


--
-- Name: troika; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE troika (
)
INHERITS (leaders);


ALTER TABLE troika OWNER TO postgres;

--
-- Name: TABLE troika; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE troika IS 'A faction.';


--
-- Name: turn_order_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE turn_order_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE turn_order_seq OWNER TO postgres;

--
-- Name: turn_order_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE turn_order_seq OWNED BY turn."order";


--
-- Name: vi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi (
)
INHERITS (deck_cards);


ALTER TABLE vi OWNER TO postgres;

--
-- Name: vi_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_actions (
)
INHERITS (vi, actions);


ALTER TABLE vi_actions OWNER TO postgres;

--
-- Name: vi_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_channeled_actions (
)
INHERITS (vi_actions, channeled);


ALTER TABLE vi_channeled_actions OWNER TO postgres;

--
-- Name: vi_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_events (
)
INHERITS (vi, events);


ALTER TABLE vi_events OWNER TO postgres;

--
-- Name: vi_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_channeled_events (
)
INHERITS (vi_events, channeled);


ALTER TABLE vi_channeled_events OWNER TO postgres;

--
-- Name: vi_followers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_followers (
)
INHERITS (vi, followers);


ALTER TABLE vi_followers OWNER TO postgres;

--
-- Name: vi_heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE vi_heroes (
)
INHERITS (vi, heroes);


ALTER TABLE vi_heroes OWNER TO postgres;

--
-- Name: wisp; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp (
)
INHERITS (deck_cards);


ALTER TABLE wisp OWNER TO postgres;

--
-- Name: wisp_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_actions (
)
INHERITS (wisp, actions);


ALTER TABLE wisp_actions OWNER TO postgres;

--
-- Name: wisp_channeled_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_channeled_actions (
)
INHERITS (wisp_actions, channeled);


ALTER TABLE wisp_channeled_actions OWNER TO postgres;

--
-- Name: wisp_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_events (
)
INHERITS (wisp, events);


ALTER TABLE wisp_events OWNER TO postgres;

--
-- Name: wisp_channeled_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_channeled_events (
)
INHERITS (wisp_events, channeled);


ALTER TABLE wisp_channeled_events OWNER TO postgres;

--
-- Name: wisp_heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_heroes (
)
INHERITS (wisp, heroes);


ALTER TABLE wisp_heroes OWNER TO postgres;

--
-- Name: wisp_channeled_heroes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_channeled_heroes (
)
INHERITS (wisp_heroes, channeled);


ALTER TABLE wisp_channeled_heroes OWNER TO postgres;

--
-- Name: wisp_continuous_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wisp_continuous_events (
)
INHERITS (wisp_events, continuous);


ALTER TABLE wisp_continuous_events OWNER TO postgres;

--
-- Name: actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: activated_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY activated_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: activities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY activities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: bast resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: bast_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: bast_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: bast_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: bast_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: bast_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: bast_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: channeled resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY channeled ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: channeled copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY channeled ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: conditional_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: conditional_card_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_card_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: conditional_lane_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_lane_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: constant_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: constant_card_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_card_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: constant_character_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_character_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: continuous resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY continuous ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: continuous copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY continuous ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: deck_characters damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: deck_characters speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: deck_characters life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters ALTER COLUMN life SET DEFAULT 1;


--
-- Name: deck_characters resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: deck_characters copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: deck_heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: deck_heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: deck_heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Name: deck_heroes resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_heroes ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: deck_heroes copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_heroes ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Name: igrath resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: igrath copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: igrath_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: igrath_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: igrath_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: igrath_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: igrath_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: igrath_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: igrath_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: igrath_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: igrath_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: leaders damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY leaders ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: leaders speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY leaders ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: leaders life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY leaders ALTER COLUMN life SET DEFAULT 1;


--
-- Name: lilith resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: lilith_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: lilith_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: lilith_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: lilith_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: nightmares damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY nightmares ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: nightmares speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY nightmares ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: nightmares life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY nightmares ALTER COLUMN life SET DEFAULT 1;


--
-- Name: partners damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY partners ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: partners speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY partners ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: partners life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY partners ALTER COLUMN life SET DEFAULT 1;


--
-- Name: ravat resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_continuous_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_continuous_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: ravat_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: ravat_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: ravat_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: ravat_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: scinter resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scinter copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scinter_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scinter_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scinter_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scinter_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scinter_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scinter_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scinter_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scinter_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scinter_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: scinter_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: scinter_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: scuttler resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scuttler copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scuttler_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: scuttler_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: scuttler_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: scuttler_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: scuttler_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: tendril resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tendril copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tendril_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tendril_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tendril_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tendril_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tendril_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tendril_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tendril_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: tendril_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: tendril_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: tinsel resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_continuous_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_continuous_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: tinsel_heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: tinsel_heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: tinsel_heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Name: tinsel_heroes resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: tinsel_heroes copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: triggered_abilities complexity; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY triggered_abilities ALTER COLUMN complexity SET DEFAULT 1;


--
-- Name: troika damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY troika ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: troika speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY troika ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: troika life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY troika ALTER COLUMN life SET DEFAULT 1;


--
-- Name: turn order; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY turn ALTER COLUMN "order" SET DEFAULT nextval('turn_order_seq'::regclass);


--
-- Name: vi resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_followers resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_followers copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_followers damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: vi_followers speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: vi_followers life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers ALTER COLUMN life SET DEFAULT 1;


--
-- Name: vi_heroes resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: vi_heroes copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: vi_heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: vi_heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: vi_heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Name: wisp resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_channeled_actions resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_actions ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_channeled_actions copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_actions ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_channeled_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_channeled_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_channeled_heroes resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_channeled_heroes copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_channeled_heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: wisp_channeled_heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: wisp_channeled_heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Name: wisp_continuous_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_continuous_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_continuous_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_continuous_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_events resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_events ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_events copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_events ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_heroes resolve_cost; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes ALTER COLUMN resolve_cost SET DEFAULT '1'::text;


--
-- Name: wisp_heroes copies; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes ALTER COLUMN copies SET DEFAULT 3;


--
-- Name: wisp_heroes damage; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes ALTER COLUMN damage SET DEFAULT 1;


--
-- Name: wisp_heroes speed; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes ALTER COLUMN speed SET DEFAULT 1;


--
-- Name: wisp_heroes life; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes ALTER COLUMN life SET DEFAULT 1;


--
-- Data for Name: abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO abilities (name, root, complexity) VALUES ('Resurrect', 'If this character would be discarded, instead they become downed.
    Downed- Put N Resurrect counters on this card.
    Start- Remove a Resurrect counter from this card, then, if it has no Resurrect counters on it, revive it.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Deadly', 'When this character deals damage to another deck character, that character gets discarded.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Resurrect', 'If this character would be discarded, instead they become downed.
    Downed- Put N Resurrect counters on this card.
    Start- Remove a Resurrect counter from this card, then, if it has no Resurrect counters on it, revive it.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Replace N', 'You may replace one card up to N times.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Stitch N', 'If this character is in your discard, you may pay N to stitch it onto a sandman in play. If you do, that sandman gains this character''s damage and life.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Maximum Value', 'Starting value, adjusted by any affects that may have altered it, i.e. +x/+x effects, damage, etc.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Owner', 'A player owns a card if they deployed it or its leader when setting up the game.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Skirmish', 'A group of one or more attacking characters and one defending character. When a skirmish ends, skirmishing characters deal damage to eachother.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Gain', 'To add or modify an ability  to a card. ', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Lose', 'To remove an ability or trait from a ${restriction}.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Deck', '20 cards that share a deck icon with their leader.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Control', 'A card''s controller is the player that can make decisions for the card. Players control all cards they put into play, unless changed by an ability, i.e. Seize or Ultimatum.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Attacker', 'Short for Attacking Player.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Current Value', 'Maxium value, adjusted by any effects that may have altered it. i.e. Taking damage, spending speed, etc.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Defender', 'Short for Defending Player.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Defending Player', 'If it isn''t a player''s turn, they are a defending player.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Downed', 'A downed card is turned sideways, and treated as though it weren''t there.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Effect', 'Anything that happens as a result of an ability.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Engaged', 'A character is engaged if they are involved in a skirmish.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Factions', 'Leaders of different factions can''t be played by the same player.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Flank Attack', 'Also known as a cross-lane attack, flanking can''t be intercepted.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Last Stand', 'During a player''s resolve phase, If they control only one hero, that hero produces an additional resolve this turn.', 1);
INSERT INTO abilities (name, root, complexity) VALUES ('Medium range', 'Uncontested- this character can flank and reinforce.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Starting Value', 'The number printed on the card.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Stat Change', 'When a character''s stat value is increased or decreased, both it''s current and maximum values are changed.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Ties', 'if more than one player would win the game, instead, the game ends in a draw.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Toughness', 'Liife that refreshes at the end of each turn.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Unengaged', 'A character that is not currently in a skirmish.', 0);
INSERT INTO abilities (name, root, complexity) VALUES ('Winning the Game', 'A player wins the game when their opponent controls no heroes with 1 or more life.', 0);


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: activated_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO activities (name, root, complexity, cost) VALUES ('Mulligan', 'set aside one''s hand and draw a new one.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Stitch', 'pay a card''s stitch cost, and attach it to a card in play.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Scrap', 'put a card from your discard pile into your scrap pile.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Refresh', 'set a trait''s current value equal to its maximum value.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Seize', 'gain control of a card until it leaves play.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Ultimatum', 'give an opponent control of this ability.', 1, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Deal', 'assign damage to one or more characters.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Discard', 'put a card into it''s owner''s discard pile from anywhere.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Move', 'put a card in another lane. The lanes don''t need to be adjacent.', 2, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Draw', 'put the top card of your deck into your hand.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Deploy', 'put it into play from hand. Deploying uses all of a character''s speed.', 0, 'a character card''s cost');
INSERT INTO activities (name, root, complexity, cost) VALUES ('Redeploy', 'move it.', 0, '1 speed from a character');
INSERT INTO activities (name, root, complexity, cost) VALUES ('Intercept', 'engage a character in a skirmish during the Interept step.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Pay', 'perform one or more activities in order to play a card or activate an ability.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Remove', 'take one or more counters off of a card.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Place', 'put one or more counters on a card.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Play', 'pay a card''s cost and perform its effects.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Reinforce', 'perform a cross-lane intercept during the Intercept step.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Search', 'look through one''s deck, usually to find a card.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Attack', 'engage a character in a skirmish during the Attack Step.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Flip', 'turn over a double-sided card.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Flank', 'perform a cross-lane attack during the Attack step.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Replace', 'discard 1 from hand, then draw 1.', 0, NULL);
INSERT INTO activities (name, root, complexity, cost) VALUES ('Revive', 'return a downed card to play with refreshed life.', 1, NULL);


--
-- Data for Name: bast; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: bast_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: bast_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO bast_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Combust', NULL, '{"Deal 2 to a character."}', NULL, 'Let''s light it up.', '1', 3);
INSERT INTO bast_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Ignite', NULL, '{"Deal 3 to a hero."}', 'Ignite can be played on leaders, partners, or deck heroes.', 'There are two kinds of things in this world- things that can catch fire, and things that are on fire.', '1', 3);


--
-- Data for Name: bast_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO bast_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Chaotic Blast', NULL, '{"Bast deals 2 to all other characters.","Nearby characters take +1 damage."}', 'Bast does not take damage.', NULL, '3', 1);
INSERT INTO bast_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Scorching Strike', NULL, '{"Deal 1 to a follower.","Draw 1."}', NULL, '...That one.', '1', 3);


--
-- Data for Name: bast_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO bast_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Searing Focus', NULL, '{"If you would draw a card, you may put a discarded, channeled Bast card into your hand instead."}', 'A discarded card is a card in your discard pile.', 'You didn''t really think I was only going to hit you once?', '3', 1);


--
-- Data for Name: bast_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO bast_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Intensify', NULL, '{"Reveal the top 4 cards of your deck. Put 2 Bast channeled cards from among them into your hand and discard the rest."}', NULL, 'Burial or cremation- and if you choose ''burial'', you picked the wrong fight.', '2', 3);
INSERT INTO bast_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Savage Melee', NULL, '{"A Troika hero deals its damage to a nearby character."}', NULL, 'You can''t put the lid back on the can.', '2', 2);
INSERT INTO bast_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Rush of Anger', NULL, '{"A character gets +3/+0 this turn."}', 'You may target any character controlled by any player.', 'You just couldn''t keep your mouth shut.', '2', 2);
INSERT INTO bast_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Relentless Fury', NULL, '{"You may move 1 Troika hero.","That hero gains 1 speed this turn."}', NULL, 'Gotta go fast.', '3', 2);


--
-- Data for Name: bold; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO bold (regex) VALUES ('affect');
INSERT INTO bold (regex) VALUES ('all');
INSERT INTO bold (regex) VALUES ('Any player');
INSERT INTO bold (regex) VALUES ('can''t');
INSERT INTO bold (regex) VALUES ('Choose');
INSERT INTO bold (regex) VALUES ('current');
INSERT INTO bold (regex) VALUES ('don''t');
INSERT INTO bold (regex) VALUES ('Each');
INSERT INTO bold (regex) VALUES ('[X:\+\-0-9]');
INSERT INTO bold (regex) VALUES ('(an)?other');
INSERT INTO bold (regex) VALUES ('damage');
INSERT INTO bold (regex) VALUES ('reveal');
INSERT INTO bold (regex) VALUES ('top');
INSERT INTO bold (regex) VALUES ('may');
INSERT INTO bold (regex) VALUES ('turn');


--
-- Data for Name: cards; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: channeled; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: characters; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: conditional_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: conditional_card_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO conditional_card_abilities (name, root, complexity, condition) VALUES ('Unique', 'If ${condition}, choose one and discard the rest.', 1, 'you control more than one copy of this card in play');
INSERT INTO conditional_card_abilities (name, root, complexity, condition) VALUES ('Nearby', '${type}s are ${name} if ${condition}', 1, 'they are in the same lane.');


--
-- Data for Name: conditional_lane_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO conditional_lane_abilities (name, root, complexity, condition) VALUES ('Adjacent', '${type}s are ${name} if ${condition}', 0, 'they share a border.');
INSERT INTO conditional_lane_abilities (name, root, complexity, condition) VALUES ('Contested', 'A ${type} is ${name} if ${condition}', 0, 'each player controls at least one character in it.');
INSERT INTO conditional_lane_abilities (name, root, complexity, condition) VALUES ('Uncontested', 'A ${type} is ${name} if ${condition}', 1, 'it is not contested.');


--
-- Data for Name: constant_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: constant_card_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: constant_character_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO constant_character_abilities (name, root, complexity, rt_text) VALUES ('Short range', '''t flank or reinforce', 1, '''t flank or reinforce');
INSERT INTO constant_character_abilities (name, root, complexity, rt_text) VALUES ('Ambush', ' be deployed any time you could play an action card.', 1, 'be deployed any time you could play an action card.');
INSERT INTO constant_character_abilities (name, root, complexity, rt_text) VALUES ('Guard', ' intercept even when attacked.', 1, 'intercept even when attacked.');
INSERT INTO constant_character_abilities (name, root, complexity, rt_text) VALUES ('Long range', ' flank and reinforce', 1, 'flank and reinforce');
INSERT INTO constant_character_abilities (name, root, complexity, rt_text) VALUES ('Stealth', ' can only be engaged by nearby characters.', 1, 'can only be engaged by nearby characters.');


--
-- Data for Name: continuous; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: deck_cards; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: deck_characters; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: deck_heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: followers; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: igrath; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: igrath_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO igrath_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('High Ground', NULL, '{"Uncontested- characters can''t deploy or redeploy to this lane."}', 'If your opponent controls nearby characters, this card does nothing.', '...You think the view from down *there* is nice?', '3', 2);


--
-- Data for Name: igrath_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO igrath_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Drastic Measures', NULL, '{"Discard a deck card from a lane, its controller draws 1."}', 'Deck heroes count as deck cards.', 'One of these things does not belong...
It''s you.', '1', 3);
INSERT INTO igrath_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Overlooked Advantage', NULL, '{"Move a character.","Draw 1."}', NULL, 'Nobody saw this coming.', '1', 2);
INSERT INTO igrath_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Coordinated Assault', NULL, '{"Characters in uncontested lanes gain 1 speed this turn."}', NULL, 'On the count of Three! One... Two...', '4', 1);


--
-- Data for Name: igrath_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO igrath_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Bushwack Squad', NULL, '{"Uncontested- +2/+1."}', NULL, 'Our chief weapon is surprise!
Surprise and guns, guns and surprise... Our two weapons are guns and surprise.
And ruthless sarcasm!
Our *three* weapons are...', '1', 3, 1, 1, 1);
INSERT INTO igrath_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Elite Troika Striker', NULL, '{Ambush.}', NULL, 'If guns are illegal, I guess my body''s a crime.', '3', 2, 3, 1, 3);
INSERT INTO igrath_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Loyal Trooper', NULL, '{"Uncontested- +2/+0."}', NULL, 'Sometimes loyalty means not asking questions.', '2', 3, 1, 1, 3);
INSERT INTO igrath_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Springtrigger Vanguard', NULL, '{"Uncontested- +2/+1."}', NULL, 'Reach out and touch someone.', '3', 2, 2, 1, 3);
INSERT INTO igrath_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Troika Tactician', NULL, '{"Uncontested- followers in this lane get +1/+1."}', 'If your opponent controls no nearby characters, your followers in this lane get +1/+1. (Including this card.)', 'Never interrupt your enemy when he is making a mistake.', '1', 2, 1, 1, 2);


--
-- Data for Name: leaders; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: lilith; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: lilith_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO lilith_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Distraction', NULL, '{"Remove a character from a skirmish.","It can''t be attacked this turn."}', 'That character takes and deals no damage this skirmish.', 'Mow! Mow! Mow!', '1', 3);
INSERT INTO lilith_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Getting Away With It', NULL, '{"Prevent the next 2 damage that would be dealt this turn."}', NULL, 'Just in the nick of time.', '1', 3);


--
-- Data for Name: lilith_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO lilith_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Unstable Lifeflow', NULL, '{"Lilith swaps life with a nearby hero."}', NULL, 'Life doesn''t end- it gives.', '3', 1);


--
-- Data for Name: lilith_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO lilith_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Devotion', NULL, '{"When a Troika character intercepts, they gain 1 speed this turn."}', NULL, 'I can take care of myself! ...Probably.', '1', 2);


--
-- Data for Name: lilith_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO lilith_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Diligent Research', NULL, '{"Draw 3."}', NULL, 'It''s almost done loading.', '2', 1);
INSERT INTO lilith_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Quick Thinking', NULL, '{"Draw 1, then Replace 2."}', NULL, 'Light on the feet is hard to beat.', '1', 2);


--
-- Data for Name: lilith_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO lilith_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Secret Bodyguard', NULL, '{"Guard, Stealth, Deadly."}', NULL, 'He can''t be here by himself- but he can make sure you''re safe.', '4', 2, 2, 1, 4);
INSERT INTO lilith_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Interfering Librarian', NULL, '{Guard.}', NULL, 'YOU ARE OUT OF ORDER.', '3', 3, 3, 1, 4);
INSERT INTO lilith_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Tower Guard', NULL, '{"Short range.",Guard.}', NULL, 'Three men to do one job...
Pays about right.', '2', 3, 1, 1, 5);


--
-- Data for Name: nightmares; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO nightmares (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Scuttler', '+1', '{"Scrap: Scuttler gains 1 resolve."}', 'You may use this ability at any time.', NULL, 0, 1, 17, '4e344f', '00f7ff', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO nightmares (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Tendril', '+2', '{"Start- Tendril gains +X/+X, where X is 7 minus his current life.","Resurrect 3."}', '', NULL, 1, 1, 7, '0d1d1c', 'b9ff00', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO nightmares (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Ravat', '+2', '{"Short range.","When another Hero becomes downed, Flip Ravat."}', NULL, NULL, 2, 1, 14, '000000', 'ff0000', '+2', 1, 3, '+4', 'Short range.
Nearby characters get -1 speed.', NULL, NULL);
INSERT INTO nightmares (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Tinsel', '+3', '{"Short range.","{C}, Ultimatum:","• You discard 2 from hand and seizes a character that was just played.","• Flip Tinsel."}', NULL, NULL, 1, 1, 10, '8f014d', 'ff01bb', '+1', 1, 3, '+3', 'Long range.
{C}: Tinsel deals 1 to a nearby character.', NULL, NULL);
INSERT INTO nightmares (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Wisp', '+1', '{"When a deck Wisp is deployed, flip Wisp.","Deck Wisps have Ambush."}', NULL, NULL, 2, 1, 14, '86cce6', '282b30', '+1', 2, 2, '-1', 'Discard a Deck Wisp: Wisp gains 1 life.
End- Flip Wisp', NULL, NULL);


--
-- Data for Name: partners; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: ravat; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: ravat_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: ravat_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Finish the Job', NULL, '{"Discard a damaged, deck character from play."}', NULL, 'I call this the time out room.
Because you''re out of time.', '1', 3);
INSERT INTO ravat_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('No Escape', NULL, '{"A nearby character loses 1 speed this turn.","If it''s your turn, move Ravat."}', NULL, 'Sometimes it''s important to stop and reflect.', '2', 3);
INSERT INTO ravat_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Terrify', NULL, '{"Nearby characters lose 2 speed this turn."}', NULL, 'Don''t worry. That''s not going to help.', '2', 2);


--
-- Data for Name: ravat_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Gruesome Display', NULL, '{"Discard one of your heroes from a lane: take an extra turn after this one."}', NULL, 'Now look what you made me do.', '2', 1);
INSERT INTO ravat_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Hunt', NULL, '{"Choose a character. ","When Ravat damages that character this turn, he gains 1 speed and gets +1/+0 this turn."}', NULL, 'Now we are friends... Don''t move.', '2', 2);


--
-- Data for Name: ravat_continuous_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_continuous_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Decaying Ground', NULL, '{"Start- Deal 1 to all nearby, non-Ravat characters.","Nearby characters gain short range."}', NULL, NULL, '2', 3);


--
-- Data for Name: ravat_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Lay Waste', NULL, '{"Start- Discard all nearby non-Ravat cards.","Nothing can enter this lane."}', 'No ability can cause a card to enter this lane under any circumstances.', 'I''m just some schmuck who enjoys his hobbies.', '5', 1);


--
-- Data for Name: ravat_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Lacerate', NULL, '{"Deal 4."}', NULL, 'Everyone smiles on the inside.
You just have to cut a little deeper.', '2', 3);


--
-- Data for Name: ravat_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO ravat_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Mokoi', NULL, '{"Short range."}', NULL, 'The ultimate truth is pain...
I''m just a humble prophet.', '3', 2, 4, 1, 5);


--
-- Data for Name: scinter; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: scinter_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO scinter_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('False Front', NULL, '{"Deal 2.","Trap- Deal 4 instead."}', NULL, '...They never look up.', '1', 3);
INSERT INTO scinter_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Double Diversion', NULL, '{"A character loses 2 speed this turn. ","Trap- Nearby characters also lose 2 speed this turn."}', NULL, NULL, '1', 2);
INSERT INTO scinter_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Jumpframe', NULL, '{"A character gets +2/+2 this turn.","Trap- That character also gains 1 speed this turn."}', NULL, 'NOBODY expects- a giant teleporting Gatling cannon!', '1', 2);
INSERT INTO scinter_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Telepad Trap', NULL, '{"Return a follower to it''s owner''s hand.","Trap- Move a Troika character."}', NULL, 'Where you need to go is where you need to be.', '2', 2);


--
-- Data for Name: scinter_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO scinter_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Chain Overreaction', NULL, '{"Reveal the top 4 cards of your deck, you may play any traps from among them without paying their costs.","Shuffle your deck.","Trap- Reveal the top 8 cards instead."}', NULL, 'It''s not paranoia if the right people die.', '2', 2);


--
-- Data for Name: scinter_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO scinter_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Master Manipulations', NULL, '{"You decide how characters intercept this turn."}', 'You may have all or none of your opponent''s characters intercept.', 'Shhh- I''m winning.', '3', 1);


--
-- Data for Name: scinter_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO scinter_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Troika Assassin', NULL, '{Ambush.,Stealth.}', NULL, '...His name is Pepito.', '3', 2, 3, 1, 3);
INSERT INTO scinter_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Indigo', NULL, '{"Short range."}', NULL, 'We ain''t here to dance.', '1', 3, 2, 2, 3);
INSERT INTO scinter_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Nainso', NULL, '{Stealth.}', NULL, 'Quick and Quiet.', '1', 3, 2, 1, 2);


--
-- Data for Name: scuttler; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: scuttler_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Corpse Hauler', NULL, '{"Stitch 2."}', NULL, 'I had that where I WANTED it!', '3', 2, 2, 1, 2);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Dust Hulk', NULL, '{"When you stitch this character, Replace 2.","Stitch 2."}', NULL, 'This one agrees with me.', '1', 3, 2, 1, 1);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Unholy Collage', NULL, '{"Deploy- Stitch all discarded sandmen onto this character.","Stitch 1."}', NULL, 'I made it myself.', '4', 1, 0, 1, 1);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Expendable Concept', NULL, '{"Deploy- Replace 2.","Stitch 2."}', NULL, 'It''s only a draft.', '1', 3, 1, 1, 1);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Ragpicker Scout', NULL, '{"Deploy- Draw 1, then Replace 2.","Stitch 1."}', NULL, 'Sit. Stay. Spy.', '1', 3, 0, 1, 1);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Sand Golem', NULL, '{"Deploy- You may Scrap X cards to discard a nearby follower with cost X or less.","Stitch 4."}', NULL, NULL, '4', 2, 3, 1, 2);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Stitchcraft Parasite', NULL, '{"Discard- Stitch this character.","Stitch 3."}', NULL, 'If you have trouble making friends, you''re not using enough material.', '1', 2, 1, 1, 1);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Limbwrencher', NULL, '{"When this character enters play, refresh a nearby Sandman''s speed.","Stitch 6."}', 'Characters with full speed can attack, intercept, and redeploy- even if they were played this turn.', 'Again?! Put that thing back.', '3', 1, 4, 1, 2);
INSERT INTO scuttler_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Spiral Minion', NULL, '{"When another sandman is discarded from a lane, Scuttler gains {1}.","Stitch 4."}', 'When a sandman that is stitched onto is discarded, everything attached to it is discarded as well.', 'All of the life-force, with none of the pesky free will.', '2', 3, 2, 1, 2);


--
-- Data for Name: slang; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO slang (name, description) VALUES ('bank', 'to not spend one or more of the resolve generated each turn, usually to save up for an expensive card.');


--
-- Data for Name: tendril; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: tendril_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tendril_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Bind', NULL, '{"A follower gets -2/-2 this turn."}', NULL, 'Going nowhere awfully fast,
You are too weak,
Your freedom past.', '1', 3);
INSERT INTO tendril_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Snuff Out', NULL, '{"A follower gets -5/-5 this turn."}', NULL, 'Die @#$%!', '2', 2);


--
-- Data for Name: tendril_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tendril_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Dark Rebirth', NULL, '{"Remove all resurrect counters from cards in each lane."}', NULL, 'Behold the spark of dark that burns,
Stirring mass from stillness grave,
the Hand of death, appalled it turns, and None now living can be saved.', '1', 1);
INSERT INTO tendril_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Consume', NULL, '{"Discard all deck cards from each lane."}', NULL, 'Lust is pale compared to hunger,
Skin to break and blood to plunder
Desire cracking jaws asunder.', '3', 1);
INSERT INTO tendril_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Fatal Blight', NULL, '{"Discard a deck card from a lane."}', '', 'Breathing faster,', '2', 2);


--
-- Data for Name: tendril_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tendril_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Toxic Entity', NULL, '{Deadly.}', NULL, 'The worm inside
We try to hide,
But twist or tun
The envy burns.', '1', 2, 1, 1, 2);
INSERT INTO tendril_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Lyrical Terror', NULL, '{"Resurrect 2."}', NULL, 'What''s a poet without an ear?', '1', 2, 1, 1, 3);
INSERT INTO tendril_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Monsterous Outgrowth', NULL, '{"Resurrect 3."}', NULL, 'What''s a monster without the fear?', '2', 2, 3, 1, 2);
INSERT INTO tendril_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Perpetual Nightmare', NULL, '{"Resurrect 3."}', NULL, 'How drab the face without a leer.', '3', 3, 3, 1, 3);
INSERT INTO tendril_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Silver-Tongued Spawn', NULL, '{"Resurrect 2."}', NULL, 'How dull the eye without a tear.', '2', 2, 2, 1, 2);


--
-- Data for Name: tinsel; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: tinsel_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Status Quo', NULL, '{"Ultimatum: Your heroes lose 2 life or get -2/-0 this turn."}', NULL, 'I''m only at the pinnacle because I''m better  than all of you.', '2', 2);


--
-- Data for Name: tinsel_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Hair Lash', NULL, '{"Tinsel Deals 2.","Draw 1."}', NULL, 'Channeled cards must be played with their leader''s resolve.', '2', 3);


--
-- Data for Name: tinsel_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Blatant Extortion', NULL, '{Ultimatum:,"• You draw 2.","• Your opponent discards 2 from hand."}', NULL, 'Perfection doesn''t leave room for compromise.', '2', 3);
INSERT INTO tinsel_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Political Influence', NULL, '{"Seize a follower in a lane."}', NULL, '''Pulling the strings'' makes it sound like they''re puppets... They''re really just tools.', '4', 2);


--
-- Data for Name: tinsel_continuous_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_continuous_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Public Figure', NULL, '{"When a follower you don''t control is discarded from play, you may search it''s owner''s deck for a card with the same name, seize it, and deploy it."}', NULL, 'Attention is power. Watch this.', '4', 1);
INSERT INTO tinsel_continuous_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Benevolent Facade', NULL, '{"Players play with the top card of their deck revealed.","Discard 1: Seize the top card of a deck this turn.","{2}: Discard the top card of a deck."}', 'Channeled cards must be played with their leader''s resolve.', 'Show me yours- I''ll show you mine.', '3', 2);


--
-- Data for Name: tinsel_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Interrogation Chamber', NULL, '{"Play this on a character, that character can''t spend speed.","{3}: Discard Prison Sentence. Any player may use this ability."}', NULL, 'You''re not pretty enough to be above the law.', '1', 3);


--
-- Data for Name: tinsel_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Deep Connections', NULL, '{"Sieze a discarded card, you may play it this turn."}', 'You must still pay all costs to play the card.', 'My people will be in touch.', '2', 2);


--
-- Data for Name: tinsel_heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tinsel_heroes (name, resolve, abilities, reminder, flavor, damage, speed, life, resolve_cost, copies) VALUES ('Stiletto', '+1', '{"When this hero deals damage, draw 1.",Unique.}', NULL, 'Never underestimate a sharp set of heels.', 2, 2, 6, '2', 2);


--
-- Data for Name: tolerances; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO tolerances (name, px) VALUES ('title', 55);
INSERT INTO tolerances (name, px) VALUES ('short', 17);
INSERT INTO tolerances (name, px) VALUES ('long', 13);
INSERT INTO tolerances (name, px) VALUES ('flavor', 80);
INSERT INTO tolerances (name, px) VALUES ('bottom', 65);


--
-- Data for Name: triggered_abilities; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: troika; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO troika (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Lilith', '+2', '{"Short range.","End- Nearby Troika heroes gain 1 life.","When Lilith becomes downed, flip her."}', NULL, NULL, 1, 1, 12, '017538', '00ff73', '+1', 1, 1, '+X', 'Long range.
Flipped- Set nearby heroes life to 1. Then, Lilith gains X life, where X is the life lost this way.', 'Lilith drains life from all nearby heroes, not just enemies.', NULL);
INSERT INTO troika (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Vi', '+2', '{"Short range.","If 2+ nearby characters attacked this turn, flip Vi."}', 'It doesn''nondeckcardst matter if the characters attacked during the the same skirmish or not.', NULL, 2, 2, 12, 'ff2190', '2c2328', '+2', 2, 2, '+2', 'End- Flip Vi.', NULL, NULL);
INSERT INTO troika (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Igrath', '+2', '{"{4}: all characters gain +1 speed this turn.","Flip Igrath."}', NULL, NULL, 2, 1, 12, '31336c', 'b6bbff', '+2', 1, 2, '+2', 'Lanes count as uncontested if you control the most characters in them.', 'Characters can flank and reinforce from uncontested lanes.', NULL);
INSERT INTO troika (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Bast', '+2', '{"Short range.","Pay 1 speed: flip Bast."}', NULL, NULL, 2, 1, 14, 'b4000c', 'ff9758', '+2', 1, 1, '+0', 'Cards that channel Bast deal +1.
Pay 1 speed: flip Bast.', NULL, NULL);
INSERT INTO troika (name, resolve, abilities, reminder, flavor, damage, speed, life, banner, indicator, resolve_b, speed_b, damage_b, life_b, short_b, reminder_b, flavor_b) VALUES ('Scinter', '+1', '{"Flipped- Move Scinter.","End- Flip Scinter."}', 'Scinter only flips at the end of your turn.', NULL, 3, 1, 15, '9c7731', 'ffc861', '+1', 1, 2, '+0', 'Stealth.
{1}: Flip Scinter.', NULL, NULL);


--
-- Data for Name: turn; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO turn ("order", name, type, definition) VALUES (1, 'Start', 'Phase', NULL);
INSERT INTO turn ("order", name, type, definition) VALUES (2, 'Resolve', 'Phase', 'The Attacking Player generates resolve on their cards.');
INSERT INTO turn ("order", name, type, definition) VALUES (3, 'Draw', 'Phase', 'The Attacking Player draws a card.');
INSERT INTO turn ("order", name, type, definition) VALUES (4, 'Replace', 'Phase', 'The Attacking Player may replace a card.');
INSERT INTO turn ("order", name, type, definition) VALUES (5, 'Play', 'Phase', 'The Attacking Player may play/deploy cards, redeploy characters, activate abilities, or start a combat.');
INSERT INTO turn ("order", name, type, definition) VALUES (6, 'Combat', 'Phase', 'A sequence of steps that allow players to engage their characters in skirmishes.');
INSERT INTO turn ("order", name, type, definition) VALUES (7, 'Attack', 'Step', 'Occurs once during combat. Allows the attacker to engage their chacters in skirmishes.');
INSERT INTO turn ("order", name, type, definition) VALUES (8, 'Intercept', 'Step', 'Occurs once during combat. Allows the defender to engage their chacters in skirmishes.');
INSERT INTO turn ("order", name, type, definition) VALUES (9, 'End', 'Phase', 'All character''s speed refreshes, as do all follower''s toughness.');


--
-- Data for Name: vi; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: vi_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Emergency Extraction', NULL, '{"A Troika hero can''t affect or be affected by other cards this turn."}', 'The character can''t take damage, deal damage, or otherwise be changed by any other card this turn. If they are in a skirmish, they are immediately withdrawn.', 'Watch Out!', '1', 3);
INSERT INTO vi_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Midrange Mk. III', NULL, '{"A character gains +2/+0 this turn.","A nearby character gains Guard this turn."}', NULL, 'When in doubt- empty your magazine.', '2', 2);


--
-- Data for Name: vi_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Hurricane Force', NULL, '{"Move a nearby character."}', 'The character must be in the same lane as Vi.', 'Reason with this.', '1', 2);


--
-- Data for Name: vi_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Boyfriend', NULL, '{"Vi gains +4/+0 and long range this turn."}', 'Long range means this character can flank and reinforce from contested lanes.', 'At least this one''s a straight shooter.', '4', 1);


--
-- Data for Name: vi_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Confrontation', NULL, '{"2 nearby characters skirmish."}', NULL, '...This is happening now.', '1', 3);


--
-- Data for Name: vi_followers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_followers (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Troika Battlebruiser', NULL, '{"Short range."}', NULL, 'Their morals are loose, but their shirts are tight.', '2', 3, 3, 1, 4);


--
-- Data for Name: vi_heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO vi_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Bobby', '+1', '{"{1}: Nearby non-Troika characters  get -1/-0 this turn.",Unique.}', NULL, 'I hate not hating you.', '3', 2, 3, 1, 9);
INSERT INTO vi_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Rumour', '+1', '{"{1}: Look at the top 3 cards of your deck. You may rearrange them.",Unique.}', NULL, 'Don''t touch his stuff.', '2', 2, 2, 1, 8);
INSERT INTO vi_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Karo', '+1', '{"Pay 1 Speed: Karo deals 2 to a nearby character. Use this ability only during your turn.",Unique.}', NULL, 'Violence isn''t the answer- but that''s okay, this isn''t a quiz!', '3', 2, 3, 1, 9);


--
-- Data for Name: wisp; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: wisp_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: wisp_channeled_actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO wisp_channeled_actions (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Tearing Swarm', NULL, '{"Deal X to a character, where X is the number of nearby Wisps."}', 'Deck Wisps count as Wisps.', 'Why are you running from me? Wisp is nice! Wisp is nice! NOW HOLD STILL.', '1', 3);


--
-- Data for Name: wisp_channeled_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO wisp_channeled_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('The Gang''s All Here', '+0', '{"Deploy up to X Wisps from your deck without paying their costs.","Shuffle your deck.","Leader Wisp loses X life."}', NULL, 'I always try to follow my hearts.', 'X', 1);


--
-- Data for Name: wisp_channeled_heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO wisp_channeled_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Deception', '+1', '{"Deploy- An opponent discards 1."}', NULL, 'Straighten up and fly right. Or left?', '1', 3, 2, 1, 4);
INSERT INTO wisp_channeled_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Anger', '+1', '{"Deploy- Anger deals 2 to a nearby non-Wisp character."}', NULL, '...Issues', '1', 3, 3, 1, 3);
INSERT INTO wisp_channeled_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Love', '+1', '{"Deploy- Each hero gains 1 life."}', NULL, 'Everyone loves Wisp.', '0', 3, 1, 1, 3);
INSERT INTO wisp_channeled_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Paranoia', '+1', '{"Deploy- Prevent the effects of an action or event card unless its controller pays {1}."}', 'This ability can not prevent continuous events from working.', 'I didn''t change my mind, I just made it go away for a little while.', '1', 2, 2, 1, 5);
INSERT INTO wisp_channeled_heroes (name, resolve, abilities, reminder, flavor, resolve_cost, copies, damage, speed, life) VALUES ('Sadness', '+2', '{"Deploy- Sadness and a nearby non-Wisp character lose 2 speed this turn."}', NULL, 'My home exploded and coins came out.', '2', 3, 2, 1, 3);


--
-- Data for Name: wisp_continuous_events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO wisp_continuous_events (name, resolve, abilities, reminder, flavor, resolve_cost, copies) VALUES ('Support Group', NULL, '{"When a Wisp is deployed, each Wisp gains X life, where X is the number nearby Wisps."}', 'Deck Wisps count as Wisps.', 'How many Wisps have you got in there?', '2', 2);


--
-- Data for Name: wisp_events; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: wisp_heroes; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: zones; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO zones (name, definition) VALUES ('Deck', 'A facedown pile of cards each player begins the game with.');
INSERT INTO zones (name, definition) VALUES ('Hand', 'A group of cards drawn from a player''s deck. Hands are private.');
INSERT INTO zones (name, definition) VALUES ('Scrap', 'A faceup pile of cards that have been removed from the game by Scuttler''s ''Scrap'' Ability.');
INSERT INTO zones (name, definition) VALUES ('Play', 'An area subdivided into 3 lanes. If a card exists in play, it must be in one of the lanes.');
INSERT INTO zones (name, definition) VALUES ('Discard', 'A faceup pile of cards that have been discarded. Each player has their own discard pile.');
INSERT INTO zones (name, definition) VALUES ('Lane', 'A division of the play area where cards can be played. Skirmish has 3 lanes.');


--
-- Name: turn_order_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('turn_order_seq', 9, true);


--
-- Name: actions actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY actions
    ADD CONSTRAINT actions_name_key UNIQUE (name);


--
-- Name: activated_abilities activated_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY activated_abilities
    ADD CONSTRAINT activated_abilities_name_key UNIQUE (name);


--
-- Name: activities activities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY activities
    ADD CONSTRAINT activities_name_key UNIQUE (name);


--
-- Name: bast_actions bast_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_actions
    ADD CONSTRAINT bast_actions_name_key UNIQUE (name);


--
-- Name: bast_channeled_actions bast_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_actions
    ADD CONSTRAINT bast_channeled_actions_name_key UNIQUE (name);


--
-- Name: bast_channeled_events bast_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_channeled_events
    ADD CONSTRAINT bast_channeled_events_name_key UNIQUE (name);


--
-- Name: bast_continuous_events bast_continuous_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_continuous_events
    ADD CONSTRAINT bast_continuous_channeled_events_name_key UNIQUE (name);


--
-- Name: bast_events bast_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast_events
    ADD CONSTRAINT bast_events_name_key UNIQUE (name);


--
-- Name: bast bast_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bast
    ADD CONSTRAINT bast_name_key UNIQUE (name);


--
-- Name: cards cards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY cards
    ADD CONSTRAINT cards_pkey PRIMARY KEY (name);


--
-- Name: channeled channeled_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY channeled
    ADD CONSTRAINT channeled_name_key UNIQUE (name);


--
-- Name: characters characters_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY characters
    ADD CONSTRAINT characters_name_key UNIQUE (name);


--
-- Name: characters characters_name_key1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY characters
    ADD CONSTRAINT characters_name_key1 UNIQUE (name);


--
-- Name: characters characters_name_key2; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY characters
    ADD CONSTRAINT characters_name_key2 UNIQUE (name);


--
-- Name: conditional_abilities conditional_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_abilities
    ADD CONSTRAINT conditional_abilities_name_key UNIQUE (name);


--
-- Name: conditional_card_abilities conditional_card_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_card_abilities
    ADD CONSTRAINT conditional_card_abilities_name_key UNIQUE (name);


--
-- Name: conditional_lane_abilities conditional_lane_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY conditional_lane_abilities
    ADD CONSTRAINT conditional_lane_abilities_name_key UNIQUE (name);


--
-- Name: constant_abilities constant_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_abilities
    ADD CONSTRAINT constant_abilities_name_key UNIQUE (name);


--
-- Name: constant_abilities constant_abilities_root_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_abilities
    ADD CONSTRAINT constant_abilities_root_key UNIQUE (rt_text);


--
-- Name: constant_card_abilities constant_card_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_card_abilities
    ADD CONSTRAINT constant_card_abilities_name_key UNIQUE (name);


--
-- Name: constant_card_abilities constant_card_abilities_root_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_card_abilities
    ADD CONSTRAINT constant_card_abilities_root_key UNIQUE (rt_text);


--
-- Name: constant_character_abilities constant_character_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_character_abilities
    ADD CONSTRAINT constant_character_abilities_name_key UNIQUE (name);


--
-- Name: constant_character_abilities constant_character_abilities_root_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY constant_character_abilities
    ADD CONSTRAINT constant_character_abilities_root_key UNIQUE (rt_text);


--
-- Name: continuous continuous_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY continuous
    ADD CONSTRAINT continuous_name_key UNIQUE (name);


--
-- Name: deck_cards deck_cards_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_cards
    ADD CONSTRAINT deck_cards_name_key UNIQUE (name);


--
-- Name: deck_characters deck_characters_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY deck_characters
    ADD CONSTRAINT deck_characters_name_key UNIQUE (name);


--
-- Name: events events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY events
    ADD CONSTRAINT events_name_key UNIQUE (name);


--
-- Name: followers followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY followers
    ADD CONSTRAINT followers_name_key UNIQUE (name);


--
-- Name: igrath_continuous_events igrath_continuous_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_continuous_events
    ADD CONSTRAINT igrath_continuous_events_name_key UNIQUE (name);


--
-- Name: igrath_events igrath_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_events
    ADD CONSTRAINT igrath_events_name_key UNIQUE (name);


--
-- Name: igrath_followers igrath_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath_followers
    ADD CONSTRAINT igrath_followers_name_key UNIQUE (name);


--
-- Name: igrath igrath_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY igrath
    ADD CONSTRAINT igrath_name_key UNIQUE (name);


--
-- Name: leaders leaders_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY leaders
    ADD CONSTRAINT leaders_name_key UNIQUE (name);


--
-- Name: lilith_actions lilith_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_actions
    ADD CONSTRAINT lilith_actions_name_key UNIQUE (name);


--
-- Name: lilith_channeled_events lilith_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_channeled_events
    ADD CONSTRAINT lilith_channeled_events_name_key UNIQUE (name);


--
-- Name: lilith_continuous_events lilith_continuous_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_continuous_events
    ADD CONSTRAINT lilith_continuous_events_name_key UNIQUE (name);


--
-- Name: lilith_events lilith_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_events
    ADD CONSTRAINT lilith_events_name_key UNIQUE (name);


--
-- Name: lilith_followers lilith_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith_followers
    ADD CONSTRAINT lilith_followers_name_key UNIQUE (name);


--
-- Name: lilith lilith_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY lilith
    ADD CONSTRAINT lilith_name_key UNIQUE (name);


--
-- Name: nightmares nightmares_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY nightmares
    ADD CONSTRAINT nightmares_name_key UNIQUE (name);


--
-- Name: partners partners_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY partners
    ADD CONSTRAINT partners_name_key UNIQUE (name);


--
-- Name: ravat_actions ravat_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_actions
    ADD CONSTRAINT ravat_actions_name_key UNIQUE (name);


--
-- Name: ravat_channeled_actions ravat_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_actions
    ADD CONSTRAINT ravat_channeled_actions_name_key UNIQUE (name);


--
-- Name: ravat_channeled_events ravat_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_channeled_events
    ADD CONSTRAINT ravat_channeled_events_name_key UNIQUE (name);


--
-- Name: ravat_continuous_channeled_events ravat_continuous_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_channeled_events
    ADD CONSTRAINT ravat_continuous_channeled_events_name_key UNIQUE (name);


--
-- Name: ravat_continuous_events ravat_continuous_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_continuous_events
    ADD CONSTRAINT ravat_continuous_events_name_key UNIQUE (name);


--
-- Name: ravat_events ravat_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_events
    ADD CONSTRAINT ravat_events_name_key UNIQUE (name);


--
-- Name: ravat_followers ravat_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat_followers
    ADD CONSTRAINT ravat_followers_name_key UNIQUE (name);


--
-- Name: ravat ravat_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY ravat
    ADD CONSTRAINT ravat_name_key UNIQUE (name);


--
-- Name: scinter_actions scinter_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_actions
    ADD CONSTRAINT scinter_actions_name_key UNIQUE (name);


--
-- Name: scinter_channeled_actions scinter_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_channeled_actions
    ADD CONSTRAINT scinter_channeled_actions_name_key UNIQUE (name);


--
-- Name: scinter_events scinter_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_events
    ADD CONSTRAINT scinter_events_name_key UNIQUE (name);


--
-- Name: scinter_followers scinter_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter_followers
    ADD CONSTRAINT scinter_followers_name_key UNIQUE (name);


--
-- Name: scinter scinter_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scinter
    ADD CONSTRAINT scinter_name_key UNIQUE (name);


--
-- Name: scuttler_followers scuttler_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler_followers
    ADD CONSTRAINT scuttler_followers_name_key UNIQUE (name);


--
-- Name: scuttler scuttler_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY scuttler
    ADD CONSTRAINT scuttler_name_key UNIQUE (name);


--
-- Name: slang slang_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY slang
    ADD CONSTRAINT slang_pkey PRIMARY KEY (name);


--
-- Name: tendril_actions tendril_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_actions
    ADD CONSTRAINT tendril_actions_name_key UNIQUE (name);


--
-- Name: tendril_events tendril_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_events
    ADD CONSTRAINT tendril_events_name_key UNIQUE (name);


--
-- Name: tendril_followers tendril_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril_followers
    ADD CONSTRAINT tendril_followers_name_key UNIQUE (name);


--
-- Name: tendril tendril_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tendril
    ADD CONSTRAINT tendril_name_key UNIQUE (name);


--
-- Name: tinsel_actions tinsel_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_actions
    ADD CONSTRAINT tinsel_actions_name_key UNIQUE (name);


--
-- Name: tinsel_channeled_actions tinsel_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_actions
    ADD CONSTRAINT tinsel_channeled_actions_name_key UNIQUE (name);


--
-- Name: tinsel_channeled_events tinsel_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_channeled_events
    ADD CONSTRAINT tinsel_channeled_events_name_key UNIQUE (name);


--
-- Name: tinsel_continuous_channeled_events tinsel_continuous_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_channeled_events
    ADD CONSTRAINT tinsel_continuous_channeled_events_name_key UNIQUE (name);


--
-- Name: tinsel_continuous_events tinsel_continuous_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_continuous_events
    ADD CONSTRAINT tinsel_continuous_events_name_key UNIQUE (name);


--
-- Name: tinsel_events tinsel_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_events
    ADD CONSTRAINT tinsel_events_name_key UNIQUE (name);


--
-- Name: tinsel_heroes tinsel_heroes_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel_heroes
    ADD CONSTRAINT tinsel_heroes_name_key UNIQUE (name);


--
-- Name: tinsel tinsel_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tinsel
    ADD CONSTRAINT tinsel_name_key UNIQUE (name);


--
-- Name: tolerances tolerances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tolerances
    ADD CONSTRAINT tolerances_pkey PRIMARY KEY (name);


--
-- Name: triggered_abilities triggered_abilities_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY triggered_abilities
    ADD CONSTRAINT triggered_abilities_name_key UNIQUE (name);


--
-- Name: troika troika_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY troika
    ADD CONSTRAINT troika_name_key UNIQUE (name);


--
-- Name: turn turn_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY turn
    ADD CONSTRAINT turn_name_key UNIQUE (name);


--
-- Name: turn turn_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY turn
    ADD CONSTRAINT turn_pkey PRIMARY KEY ("order");


--
-- Name: bold unique_restraint; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY bold
    ADD CONSTRAINT unique_restraint UNIQUE (regex);


--
-- Name: vi_actions vi_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_actions
    ADD CONSTRAINT vi_actions_name_key UNIQUE (name);


--
-- Name: vi_channeled_actions vi_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_actions
    ADD CONSTRAINT vi_channeled_actions_name_key UNIQUE (name);


--
-- Name: vi_channeled_events vi_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_channeled_events
    ADD CONSTRAINT vi_channeled_events_name_key UNIQUE (name);


--
-- Name: vi_events vi_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_events
    ADD CONSTRAINT vi_events_name_key UNIQUE (name);


--
-- Name: vi_followers vi_followers_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_followers
    ADD CONSTRAINT vi_followers_name_key UNIQUE (name);


--
-- Name: vi_heroes vi_heroes_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi_heroes
    ADD CONSTRAINT vi_heroes_name_key UNIQUE (name);


--
-- Name: vi vi_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY vi
    ADD CONSTRAINT vi_name_key UNIQUE (name);


--
-- Name: wisp_actions wisp_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_actions
    ADD CONSTRAINT wisp_actions_name_key UNIQUE (name);


--
-- Name: wisp_channeled_actions wisp_channeled_actions_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_actions
    ADD CONSTRAINT wisp_channeled_actions_name_key UNIQUE (name);


--
-- Name: wisp_channeled_events wisp_channeled_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_events
    ADD CONSTRAINT wisp_channeled_events_name_key UNIQUE (name);


--
-- Name: wisp_channeled_heroes wisp_channeled_heroes_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_channeled_heroes
    ADD CONSTRAINT wisp_channeled_heroes_name_key UNIQUE (name);


--
-- Name: wisp_continuous_events wisp_continuous_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_continuous_events
    ADD CONSTRAINT wisp_continuous_events_name_key UNIQUE (name);


--
-- Name: wisp_events wisp_events_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_events
    ADD CONSTRAINT wisp_events_name_key UNIQUE (name);


--
-- Name: wisp_heroes wisp_heroes_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp_heroes
    ADD CONSTRAINT wisp_heroes_name_key UNIQUE (name);


--
-- Name: wisp wisp_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wisp
    ADD CONSTRAINT wisp_name_key UNIQUE (name);


--
-- Name: zones zones_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY zones
    ADD CONSTRAINT zones_pkey PRIMARY KEY (name);


--
-- Name: bast bast_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE bast_insert_default AS
    ON INSERT TO bast DO INSTEAD NOTHING;


--
-- Name: igrath igrath_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE igrath_insert_default AS
    ON INSERT TO igrath DO INSTEAD NOTHING;


--
-- Name: lilith lilith_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE lilith_insert_default AS
    ON INSERT TO lilith DO INSTEAD NOTHING;


--
-- Name: ravat ravat_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE ravat_insert_default AS
    ON INSERT TO ravat DO INSTEAD NOTHING;


--
-- Name: scinter scinter_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE scinter_insert_default AS
    ON INSERT TO scinter DO INSTEAD NOTHING;


--
-- Name: scuttler scuttler_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE scuttler_insert_default AS
    ON INSERT TO scuttler DO INSTEAD NOTHING;


--
-- Name: tendril tendril_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE tendril_insert_default AS
    ON INSERT TO tendril DO INSTEAD NOTHING;


--
-- Name: tinsel tinsel_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE tinsel_insert_default AS
    ON INSERT TO tinsel DO INSTEAD NOTHING;


--
-- Name: vi vi_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE vi_insert_default AS
    ON INSERT TO vi DO INSTEAD NOTHING;


--
-- Name: wisp wisp_insert_default; Type: RULE; Schema: public; Owner: postgres
--

CREATE RULE wisp_insert_default AS
    ON INSERT TO wisp DO INSTEAD NOTHING;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;
GRANT ALL ON SCHEMA public TO postgres WITH GRANT OPTION;


--
-- Name: TABLE abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE abilities TO guest;


--
-- Name: TABLE cards; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE cards TO guest;


--
-- Name: TABLE deck_cards; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE deck_cards TO guest;


--
-- Name: TABLE abilities_with_type; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE abilities_with_type TO guest;


--
-- Name: TABLE ability_types; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ability_types TO guest;


--
-- Name: TABLE events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE events TO guest;


--
-- Name: TABLE actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE actions TO guest;


--
-- Name: TABLE activated_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE activated_abilities TO guest;


--
-- Name: TABLE activities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE activities TO guest;


--
-- Name: TABLE bast; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast TO guest;


--
-- Name: TABLE bast_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast_actions TO guest;


--
-- Name: TABLE channeled; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE channeled TO guest;


--
-- Name: TABLE bast_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast_channeled_actions TO guest;


--
-- Name: TABLE bast_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast_events TO guest;


--
-- Name: TABLE bast_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast_channeled_events TO guest;


--
-- Name: TABLE continuous; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE continuous TO guest;


--
-- Name: TABLE bast_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bast_continuous_events TO guest;


--
-- Name: TABLE bold; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE bold TO guest;


--
-- Name: TABLE inheritence; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE inheritence TO guest;


--
-- Name: TABLE card_traits; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE card_traits TO guest;


--
-- Name: TABLE characters; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE characters TO guest;


--
-- Name: TABLE heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE heroes TO guest;


--
-- Name: TABLE leaders; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE leaders TO guest;


--
-- Name: TABLE zones; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE zones TO guest;


--
-- Name: TABLE card_types; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE card_types TO guest;


--
-- Name: TABLE completed; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE completed TO guest;


--
-- Name: TABLE conditional_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE conditional_abilities TO guest;


--
-- Name: TABLE conditional_card_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE conditional_card_abilities TO guest;


--
-- Name: TABLE conditional_lane_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE conditional_lane_abilities TO guest;


--
-- Name: TABLE constant_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE constant_abilities TO guest;


--
-- Name: TABLE constant_card_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE constant_card_abilities TO guest;


--
-- Name: TABLE constant_character_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE constant_character_abilities TO guest;


--
-- Name: TABLE deck_characters; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE deck_characters TO guest;


--
-- Name: TABLE deck_heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE deck_heroes TO guest;


--
-- Name: TABLE followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE followers TO guest;


--
-- Name: TABLE turn; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE turn TO guest;


--
-- Name: TABLE glossary; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE glossary TO guest;


--
-- Name: TABLE igrath; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE igrath TO guest;


--
-- Name: TABLE igrath_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE igrath_events TO guest;


--
-- Name: TABLE igrath_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE igrath_continuous_events TO guest;


--
-- Name: TABLE igrath_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE igrath_followers TO guest;


--
-- Name: TABLE lilith; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith TO guest;


--
-- Name: TABLE lilith_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith_actions TO guest;


--
-- Name: TABLE lilith_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith_events TO guest;


--
-- Name: TABLE lilith_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith_channeled_events TO guest;


--
-- Name: TABLE lilith_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith_continuous_events TO guest;


--
-- Name: TABLE lilith_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE lilith_followers TO guest;


--
-- Name: TABLE nightmares; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE nightmares TO guest;


--
-- Name: TABLE partners; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE partners TO guest;


--
-- Name: TABLE ravat; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat TO guest;


--
-- Name: TABLE ravat_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_actions TO guest;


--
-- Name: TABLE ravat_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_channeled_actions TO guest;


--
-- Name: TABLE ravat_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_events TO guest;


--
-- Name: TABLE ravat_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_channeled_events TO guest;


--
-- Name: TABLE ravat_continuous_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_continuous_channeled_events TO guest;


--
-- Name: TABLE ravat_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_continuous_events TO guest;


--
-- Name: TABLE ravat_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE ravat_followers TO guest;


--
-- Name: TABLE scinter; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scinter TO guest;


--
-- Name: TABLE scinter_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scinter_actions TO guest;


--
-- Name: TABLE scinter_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scinter_channeled_actions TO guest;


--
-- Name: TABLE scinter_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scinter_events TO guest;


--
-- Name: TABLE scinter_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scinter_followers TO guest;


--
-- Name: TABLE scuttler; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scuttler TO guest;


--
-- Name: TABLE scuttler_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE scuttler_followers TO guest;


--
-- Name: TABLE slang; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE slang TO guest;


--
-- Name: TABLE tendril; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tendril TO guest;


--
-- Name: TABLE tendril_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tendril_actions TO guest;


--
-- Name: TABLE tendril_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tendril_events TO guest;


--
-- Name: TABLE tendril_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tendril_followers TO guest;


--
-- Name: TABLE tinsel; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel TO guest;


--
-- Name: TABLE tinsel_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_actions TO guest;


--
-- Name: TABLE tinsel_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_channeled_actions TO guest;


--
-- Name: TABLE tinsel_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_events TO guest;


--
-- Name: TABLE tinsel_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_channeled_events TO guest;


--
-- Name: TABLE tinsel_continuous_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_continuous_channeled_events TO guest;


--
-- Name: TABLE tinsel_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_continuous_events TO guest;


--
-- Name: TABLE tinsel_heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tinsel_heroes TO guest;


--
-- Name: TABLE tolerances; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE tolerances TO guest;


--
-- Name: TABLE triggered_abilities; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE triggered_abilities TO guest;


--
-- Name: TABLE troika; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE troika TO guest;


--
-- Name: TABLE vi; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi TO guest;


--
-- Name: TABLE vi_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_actions TO guest;


--
-- Name: TABLE vi_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_channeled_actions TO guest;


--
-- Name: TABLE vi_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_events TO guest;


--
-- Name: TABLE vi_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_channeled_events TO guest;


--
-- Name: TABLE vi_followers; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_followers TO guest;


--
-- Name: TABLE vi_heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE vi_heroes TO guest;


--
-- Name: TABLE wisp; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp TO guest;


--
-- Name: TABLE wisp_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_actions TO guest;


--
-- Name: TABLE wisp_channeled_actions; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_channeled_actions TO guest;


--
-- Name: TABLE wisp_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_events TO guest;


--
-- Name: TABLE wisp_channeled_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_channeled_events TO guest;


--
-- Name: TABLE wisp_heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_heroes TO guest;


--
-- Name: TABLE wisp_channeled_heroes; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_channeled_heroes TO guest;


--
-- Name: TABLE wisp_continuous_events; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT ON TABLE wisp_continuous_events TO guest;


--
-- PostgreSQL database dump complete
--

