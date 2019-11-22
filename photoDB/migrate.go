package photoDB

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/apierr"
	migrate "github.com/rubenv/sql-migrate"
)

// TODO: write schema
// TODO: formalize migrations a little bit more
// TODO: run certain migrations only when prod mode
// TODO: handle db nulls with special type 	"gopkg.in/guregu/null.v3"

// migrateDB migrates the photos DB, and returns error if failed
func migrateDB(cx *Config) error {
	errTag := "migrateDB"

	// sum up migrations in a slice
	// TODO: get migrations in from somewhere else
	migrations := &migrate.MemoryMigrationSource{
		Migrations: migrations,
	}

	// apply migrations
	n, err := migrate.Exec(dbx.DB, cx.ConnType, migrations, migrate.Up)
	if err != nil {
		return apierr.Errorf(err, errTag, "applying migrations")
	}

	// once migrated, the migration will not be reapplied :)
	fmt.Printf("db migrated with %d migrations\n", n)

	return nil
}

// migrations is a list of database migrations
// in- memory migrations will be built into the static binary
// "Down" migrations are omitted
// this process should only ever move forward, never backwards
// even if something is to be dropped, it should be included in a NEW migration
var migrations = []*migrate.Migration{
	&migrate.Migration{
		Id: "2019-000-create-table-testy",
		Up: []string{`
CREATE TABLE IF NOT EXISTS testy(  
	test_id INT NOT NULL AUTO_INCREMENT,  
	test_name VARCHAR(100) NOT NULL,  
	PRIMARY KEY ( test_id )  
) 
		`},
	},
	&migrate.Migration{
		Id: "2019-001-create-table-album",
		Up: []string{`
CREATE TABLE IF NOT EXISTS album (
	album_id BIGINT NOT NULL AUTO_INCREMENT
		, PRIMARY KEY (album_id)
	, album_title VARCHAR(100) NOT NULL
		, UNIQUE KEY(album_title)
	, album_description VARCHAR(300) NOT NULL
	, album_key VARCHAR(100) NOT NULL
		, UNIQUE KEY(album_key)
	, album_photo_src VARCHAR(4000) NOT NULL
)
		`},
	},
	&migrate.Migration{
		Id: "2019-002-populate-table-album",
		Up: []string{`
INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
	'Riding the SAM Shortline Train 1'
	, 'From Cordele, GA to Plains, GA 1'
	, 'sam-shortline'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg 1'
)`, `
INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
	'Riding the SAM Shortline Train 2'
	, 'From Cordele, GA to Plains, GA 2'
	, 'sam-shortline 2'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg 2'
)
		`},
	},
	&migrate.Migration{
		Id: "2019-003-create-table-photo",
		Up: []string{`
CREATE TABLE IF NOT EXISTS photo (
	photo_id BIGINT NOT NULL AUTO_INCREMENT
		, PRIMARY KEY (photo_id)
	, album_id BIGINT NOT NULL
	, photo_title VARCHAR(100) NOT NULL
	, photo_description VARCHAR(300) NOT NULL
	, photo_src VARCHAR(4000) NOT NULL
)
			`},
	},
	&migrate.Migration{
		Id: "2019-004-populate-table-photo",
		Up: []string{`
INSERT INTO photo (album_id, photo_title, photo_description, photo_src) VALUES (
	1
	, 'So much fun!'
	, 'Look at that smile'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg'
)
			`},
	},
}

// `
// import { IPhoto, IPhotoGroup, IPhotoAlbum } from './photo-albums.types';

// export const photoAlbumsData: IPhotoAlbum[] = [
//   // {
//   //   title: 'Hiking with Grandy & Papa',
//   //   routeKey: 'april-2019-hiking',
//   //   coverPhotoSrc: 'april-2019-15.jpg',
//   //   photoGroups: [
//   //     {
//   //       title: 'McDaniel Farm Park in Duluth, GA',
//   //       photos: [
//   //         { src: 'april-2019-1.jpg' }, // flipped upside down - go fix
//   //         { src: 'april-2019-2.jpg' },
//   //         { src: 'april-2019-3.jpg' },
//   //         { src: 'april-2019-6.jpg' },
//   //         { src: 'april-2019-7.jpg' },
//   //         { src: 'april-2019-9.jpg' },
//   //         { src: 'april-2019-10.jpg' },
//   //         { src: 'april-2019-14.jpg' },
//   //         { src: 'april-2019-15.jpg' },
//   //         { src: 'april-2019-17.jpg' },
//   //         { src: 'april-2019-19.jpg' },
//   //         { src: 'april-2019-20.jpg' },
//   //         { src: 'april-2019-21.jpg' },
//   //         { src: 'april-2019-22.jpg' },
//   //         { src: 'april-2019-23.jpg' },
//   //       ]
//   //     },
//   //   ]
//   // },
//   {
//     title: 'Hiking with Daddy',
//     routeKey: 'feb-2019-hiking',
//     coverPhotoSrc: 'feb-2019 (55).JPG',
//     photoGroups: [
//       {
//         title: 'Sweetwater Creek State Park',
//         photos: ((): IPhoto[] => {
//           const photos: IPhoto[] = [];
//             photos.push({ src: `feb-2019.JPG` });
//             for (var i = 1; i <= 63; i++) {
//             photos.push({ src: `feb-2019 (${i}).JPG` });
//           }
//           return photos;
//         })()
//       },
//     ]
//   },
//   {
//     title: 'Riding the SAM Shortline Train',
//     routeKey: 'sam-shortline',
//     coverPhotoSrc: 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg',
//     photoGroups: [
//       {
//         title: 'Cordele to Plains, GA',
//         photos: [
//           { src: 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg' },
//           { src: 'sam-shortline-candler-pointing-1.jpg' },
//           { src: 'sam-shortline-candler-pointing-2.jpg' },
//           { src: 'sam-shortline-grandy-1.jpg' },
//           { src: 'sam-shortline-candler-grandy-1.jpg' },
//           { src: 'sam-shortline-candler-grandy-2.jpg' },
//           { src: 'sam-shortline-candler-grandy-daddy-1.jpg' },
//           { src: 'sam-shortline-candler-grandy-daddy-2.jpg' },
//           { src: 'sam-shortline-freight-cars-1.jpg' },
//           { src: 'sam-shortline-freight-cars-2.jpg' },
//           { src: 'sam-shortline-papa-1.jpg' },
//           { src: 'sam-shortline-papa-2.jpg' },
//           { src: 'sam-shortline-candler-papa-1.jpg' },
//           { src: 'sam-shortline-candler-papa-candy-cane-bump.jpg' },
//           { src: 'sam-shortline-candler-papa-grandy-walking.jpg' },
//           { src: 'sam-shortline-candler-grandy-papa-railroad-crossing.jpg' },
//           { src: 'sam-shortline-candler-grandy-papa-with-train-1.jpg' },
//           { src: 'sam-shortline-candler-grandy-papa-with-train-2.jpg' },
//           { src: 'sam-shortline-engine-in-plains.jpg' },
//           { src: 'sam-shortline-candler-papa-grandy-jimmy-carter-peanut.jpg' },
//           { src: 'sam-shortline-candler-desk-1.jpg' },
//           { src: 'sam-shortline-candler-desk-2.jpg' },
//           { src: 'sam-shortline-candler-aisle-1.jpg' },
//           { src: 'sam-shortline-candler-aisle-2.jpg' },
//           { src: 'sam-shortline-candler-hands-out.jpg' },
//           { src: 'sam-shortline-candler-map-up.jpg' },
//           { src: 'sam-shortline-candler-map-down.jpg' },
//           { src: 'sam-shortline-candler-conductor.jpg' },
//           { src: 'sam-shortline-candler-ruler-2.jpg' },
//           { src: 'sam-shortline-candler-ruler-3.jpg' },
//           { src: 'sam-shortline-candler-ruler-4.jpg' },
//           { src: 'sam-shortline-candler-sticker-on-mouth.jpg' },
//           { src: 'sam-shortline-candler-sticker-on-grandy-mouth-1.jpg' },
//           { src: 'sam-shortline-candler-sticker-on-grandy-mouth-2.jpg' },
//           { src: 'sam-shortline-candler-daddy-looking-out-window.jpg' },
//           { src: 'sam-shortline-candler-window-staring-away.jpg' },
//           { src: 'sam-shortline-candler-window-funny-face.jpg' },
//           { description: 'The freight train broke up ahead!', src: 'sam-shortline-freight-train-broken.jpg' },
//           { src: 'sam-shortline-candler-fist-pump.jpg' },
//         ]
//       },
//     ]
//   },
//   {
//     title: 'Thanksgiving at Grandy\'s House',
//     routeKey: 'thanksgiving-2018',
//     coverPhotoSrc: 'thanksgiving-2018-frank-papa-bev-1.jpg',
//     photoGroups: [
//       {
//         title: 'Thanksgiving 2018:',
//         photos: [
//           { src: 'thanksgiving-2018-alex-1.jpg' },
//           { src: 'thanksgiving-2018-alex-2.jpg' },
//           { src: 'thanksgiving-2018-bev-hannah-1.jpg' },
//           { src: 'thanksgiving-2018-bev-sam-1.jpg' },
//           { src: 'thanksgiving-2018-bev.jpg' },
//           { src: 'thanksgiving-2018-caleb-hannah-1.jpg' },
//           { src: 'thanksgiving-2018-caleb-hannah-alex-1.jpg' },
//           { src: 'thanksgiving-2018-caleb-hannah-alex-2.jpg' },
//           { src: 'thanksgiving-2018-candler-1.jpg' },
//           { src: 'thanksgiving-2018-candler-2.jpg' },
//           { src: 'thanksgiving-2018-candler-3.jpg' },
//           { src: 'thanksgiving-2018-candler-alex-daddy-grandy-1.jpg' },
//           { src: 'thanksgiving-2018-candler-alex-hannah-tickling-alex.jpg' },
//           { src: 'thanksgiving-2018-candler-alex-hannah-tickling-candler.jpg' },
//           { src: 'thanksgiving-2018-candler-alex.jpg' },
//           { src: 'thanksgiving-2018-candler-grandy-papa-1.jpg' },
//           { src: 'thanksgiving-2018-candler-grandy-papa-2.jpg' },
//           { src: 'thanksgiving-2018-candler-papa-ted-1.jpg' },
//           { src: 'thanksgiving-2018-candler-papa-ted-2.jpg' },
//           { src: 'thanksgiving-2018-dan-brandy-1.jpg' },
//           { src: 'thanksgiving-2018-dan-brandy-2.jpg' },
//           { src: 'thanksgiving-2018-frank-papa-bev-1.jpg' },
//           { src: 'thanksgiving-2018-grandy-bev-1.jpg' },
//           { src: 'thanksgiving-2018-grandy-bev-2.jpg' },
//           { src: 'thanksgiving-2018-grandy-candler-alex.jpg' },
//           { src: 'thanksgiving-2018-joseph-alex-1.jpg' },
//           { src: 'thanksgiving-2018-mom-brandy.jpg' },
//           { src: 'thanksgiving-2018-papa-ted-1.jpg' },
//           { src: 'thanksgiving-2018-sam-hannah-1.jpg' },
//           { src: 'thanksgiving-2018-sam-jennifer-caleb-hannah-1.jpg' },
//           { src: 'thanksgiving-2018-sam-jennifer-caleb-hannah-2.jpg' },
//           { src: 'thanksgiving-2018-sam-jennifer.jpg' },
//         ]
//       },

//     ]
//   },
//   {
//     title: 'Candler Playing Around',
//     routeKey: 'candler',
//     coverPhotoSrc: 'candler-bathtub-trains.jpg',
//     photoGroups: [
//       {
//         title: 'Candler:',
//         photos: [
//           { src: 'candler-orange-at-grandys.jpg' },
//           { src: 'candler-umbrella-deer.jpg' },
//           { src: 'candler-and-josh.jpg' },
//           { src: 'candler-bridge-fist-up.jpg' },
//           { src: 'candler-bridge-pose-1.jpg' },
//           { src: 'candler-bridge-pose-2.jpg' },
//           { src: 'candler-bridge-pose-3.jpg' },
//           { src: 'brownie-monster.jpg' },
//           { src: 'candler-bathtub-trains.jpg' },
//           { src: 'candler-glasses-laughing.jpg' },
//           { src: 'candler-piano-standing.jpg' },
//           { src: 'candler-piano-sitting-1.jpg' },
//           { src: 'candler-piano-snoopy.jpg' },
//           { src: 'candler-playing-trains-1.jpg' },
//           { src: 'candler-trains-knex-bridge.jpg' },
//           { src: 'candler-trains-on-floor.jpg' },
//           { src: 'candler-train-sylvester-1.jpg' },
//           { src: 'candler-train-sylvester-2.jpg' },
//           { src: 'candler-golf-range.jpg' },
//           { src: 'candler-playing-trains-1.jpg' },
//           { src: 'candler-snoopy-breakfast-casserole.jpg' },
//         ]
//       },
//     ]
//   },
//   {
//     title: 'Grandy and Papa!',
//     routeKey: 'grandy-papa',
//     coverPhotoSrc: 'grandy-and-papa.jpg',
//     photoGroups: [
//       {
//         title: 'Grandy and Papa:',
//         photos: [
//           { src: 'candler-grandy-blowing-bubbles.jpg' },
//           { src: 'candler-grandy-train-marietta.jpg' },
//           { src: 'candler-grandy-cracker-barrel-1.jpg' },
//           { src: 'candler-grandy-cracker-barrel-2.jpg' },
//           { src: 'candler-grandy-swing.jpg' },
//           { src: 'candler-papa-train-agrirama.jpg' },
//           { src: 'candler-papa-train-marietta.jpg' },
//           { src: 'grandy-and-papa.jpg' },
//           { src: 'candler-grandy-papas-house.jpg' },
//           { src: 'candler-papa-playing.jpg' },
//           { src: 'christmas-2017-candler-grandly-papa.jpg' },
//         ]
//       },
//     ]
//   },
//   {
//     title: 'With Amy & Family',
//     routeKey: 'amy-and-family',
//     coverPhotoSrc: 'eva-tina-ngan-thanh.jpg',
//     photoGroups: [
//       {
//         photos: [
//           { src: 'amy-xinh-dep-1.jpg' },
//           { src: 'amy-xinh-dep-2.jpg' },
//           { src: 'candler-and-ngan.jpg' },
//           { src: 'candler-ngan-thanh-ice-cream.jpg' },
//           { src: 'eva-ngan-amy-candler-uncle.jpg' },
//           { src: 'candler-ngan-thanh-pool.jpg' },
//           { src: 'eva-and-uncle.jpg' },
//           { src: 'eva-tina-ngan-thanh.jpg' },
//           { src: 'eva-and-grandfather-1.jpg' },
//           { src: 'eva-and-grandfather-2.jpg' },
//         ]
//       },
//     ]
//   },
//   {
//     title: 'Random Pictures',
//     routeKey: 'random',
//     coverPhotoSrc: 'eclipse-2017-2.jpg',
//     photoGroups: [
//       {
//         // title: 'Random Family:',
//         photos: [
//           { description: 'Papa Ted @ Ed\'s Truckstop', src: 'papa-ted-eds-truck-stop.jpg' },
//         ]
//       },
//       {
//         // title: 'Trains:',
//         photos: [
//           { description: 'Near Marietta Square', src: 'train-marietta.jpg' },
//         ]
//       },
//       {
//         // title: 'Eclipse 2017:',
//         description: 'Eclipse 2017:',
//         photos: [
//           { src: 'eclipse-2017-1.jpg' },
//           { src: 'eclipse-2017-2.jpg' },
//         ]
//       },
//     ]
//   },
//   {
//     title: 'Christmas 2016',
//     routeKey: 'christmas-2016',
//     coverPhotoSrc: 'christmas-2016-75.jpg',
//     photoGroups: [
//       {
//         photos: ((): IPhoto[] => {
//           const photos: IPhoto[] = [];
//           for (var i = 2; i <= 130; i++) {
//             photos.push({ src: `christmas-2016-${i}.jpg` });
//           }
//           return photos;
//         })()
//       },
//     ]
//   },
//   {
//     title: 'Thanksgiving & Random 2016',
//     routeKey: 'thanksgiving-2016',
//     coverPhotoSrc: 'thanksgiving-2016-55.jpg',
//     photoGroups: [
//       {
//         photos: ((): IPhoto[] => {
//           const numbers: number[] = [
//             2, 4, 5, 10, 13, 14, 15, 18, 19, 20,
//             22, 23, 25, 26, 30, 31, 33, 35, 37,
//             43, 44, 49, 50, 55
//           ];
//           const photos: IPhoto[] = [];
//           for (let n in numbers) {
//             photos.push({ src: `thanksgiving-2016-${numbers[n]}.jpg` });
//           }
//           return photos;
//         })()
//       },
//     ]
//   },
// ];

// `
