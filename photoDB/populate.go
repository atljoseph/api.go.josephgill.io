package photoDB

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// PopulateDB populates the DB with data using the photoDB business logic
func PopulateDB() error {
	funcTag := "PopulateDB"

	pkgLog.WithFunc(funcTag).WithMessage("populate initial data").Info()

	// create a transaction
	txo, err := NewTxO("Test User")
	if err != nil {
		return apierr.Errorf(err, funcTag, "open db transaction")
	}

	// TODO populate initial data with populateDB instead of migrate
	err = populatePhotos(txo)
	if err = txo.RollbackOnError(err); err != nil {
		return apierr.Errorf(err, funcTag, "error populating db")
	}

	// commit transaction
	err = txo.Commit()
	if err != nil {
		return apierr.Errorf(err, funcTag, "commit db transaction")
	}

	return nil
}

// populatePhotos will fail when trying to insert any album with duplicate key
func populatePhotos(txo *TxO) error {
	funcTag := "populatePhotos"
	funcLog := pkgLog.WithFunc(funcTag)

	funcLog.WithMessage("inside db transaction").Info()

	// build albums
	funcLog.WithMessage("building album objects for insertion").Info()
	albums := []*Album{
		&Album{
			Title:       "Hiking with Grandy & Papa",
			Description: "McDaniel Farm Park in Duluth, GA",
			Key:         "april-2019-hiking",
			CoverPhoto:  "april-2019-15.jpg",
			Photos: []*Photo{
				&Photo{Src: "april-2019-1.jpg"}, // flipped upside down - go fix
				&Photo{Src: "april-2019-2.jpg"},
				&Photo{Src: "april-2019-3.jpg"},
				&Photo{Src: "april-2019-6.jpg"},
				&Photo{Src: "april-2019-7.jpg"},
				&Photo{Src: "april-2019-9.jpg"},
				&Photo{Src: "april-2019-10.jpg"},
				&Photo{Src: "april-2019-14.jpg"},
				&Photo{Src: "april-2019-15.jpg"},
				&Photo{Src: "april-2019-17.jpg"},
				&Photo{Src: "april-2019-19.jpg"},
				&Photo{Src: "april-2019-20.jpg"},
				&Photo{Src: "april-2019-21.jpg"},
				&Photo{Src: "april-2019-22.jpg"},
				&Photo{Src: "april-2019-23.jpg"},
			},
		},
		&Album{
			Title:       "Hiking with Daddy",
			Description: "Sweetwater Creek State Park",
			Key:         "feb-2019-hiking",
			CoverPhoto:  "feb-2019 (55).JPG",
			Photos: func() []*Photo {
				ps := []*Photo{}
				ps = append(ps, &Photo{Src: "feb-2019.JPG"})
				for i := 0; i < 63; i++ {
					s := fmt.Sprintf("feb-2019 (%d).JPG", i)
					ps = append(ps, &Photo{Src: s})
				}
				return ps
			}(),
		},
		&Album{
			Title:       "Riding the SAM Shortline Train",
			Description: "",
			Key:         "sam-shortline",
			CoverPhoto:  "sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg",
			Photos: []*Photo{
				&Photo{Src: "sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg"},
				&Photo{Src: "sam-shortline-candler-pointing-1.jpg"},
				&Photo{Src: "sam-shortline-candler-pointing-2.jpg"},
				&Photo{Src: "sam-shortline-grandy-1.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-1.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-2.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-daddy-1.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-daddy-2.jpg"},
				&Photo{Src: "sam-shortline-freight-cars-1.jpg"},
				&Photo{Src: "sam-shortline-freight-cars-2.jpg"},
				&Photo{Src: "sam-shortline-papa-1.jpg"},
				&Photo{Src: "sam-shortline-papa-2.jpg"},
				&Photo{Src: "sam-shortline-candler-papa-1.jpg"},
				&Photo{Src: "sam-shortline-candler-papa-candy-cane-bump.jpg"},
				&Photo{Src: "sam-shortline-candler-papa-grandy-walking.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-papa-railroad-crossing.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-papa-with-train-1.jpg"},
				&Photo{Src: "sam-shortline-candler-grandy-papa-with-train-2.jpg"},
				&Photo{Src: "sam-shortline-engine-in-plains.jpg"},
				&Photo{Src: "sam-shortline-candler-papa-grandy-jimmy-carter-peanut.jpg"},
				&Photo{Src: "sam-shortline-candler-desk-1.jpg"},
				&Photo{Src: "sam-shortline-candler-desk-2.jpg"},
				&Photo{Src: "sam-shortline-candler-aisle-1.jpg"},
				&Photo{Src: "sam-shortline-candler-aisle-2.jpg"},
				&Photo{Src: "sam-shortline-candler-hands-out.jpg"},
				&Photo{Src: "sam-shortline-candler-map-up.jpg"},
				&Photo{Src: "sam-shortline-candler-map-down.jpg"},
				&Photo{Src: "sam-shortline-candler-conductor.jpg"},
				&Photo{Src: "sam-shortline-candler-ruler-2.jpg"},
				&Photo{Src: "sam-shortline-candler-ruler-3.jpg"},
				&Photo{Src: "sam-shortline-candler-ruler-4.jpg"},
				&Photo{Src: "sam-shortline-candler-sticker-on-mouth.jpg"},
				&Photo{Src: "sam-shortline-candler-sticker-on-grandy-mouth-1.jpg"},
				&Photo{Src: "sam-shortline-candler-sticker-on-grandy-mouth-2.jpg"},
				&Photo{Src: "sam-shortline-candler-daddy-looking-out-window.jpg"},
				&Photo{Src: "sam-shortline-candler-window-staring-away.jpg"},
				&Photo{Src: "sam-shortline-candler-window-funny-face.jpg"},
				&Photo{Description: "The freight train broke up ahead!",
					Src: "sam-shortline-freight-train-broken.jpg"},
				&Photo{Src: "sam-shortline-candler-fist-pump.jpg"},
			},
		},
		&Album{
			Title:       "Thanksgiving at Grandy's House",
			Description: "",
			Key:         "thanksgiving-2018",
			CoverPhoto:  "thanksgiving-2018-frank-papa-bev-1.jpg",
			Photos: []*Photo{
				&Photo{Src: "thanksgiving-2018-alex-1.jpg"},
				&Photo{Src: "thanksgiving-2018-alex-2.jpg"},
				&Photo{Src: "thanksgiving-2018-bev-hannah-1.jpg"},
				&Photo{Src: "thanksgiving-2018-bev-sam-1.jpg"},
				&Photo{Src: "thanksgiving-2018-bev.jpg"},
				&Photo{Src: "thanksgiving-2018-caleb-hannah-1.jpg"},
				&Photo{Src: "thanksgiving-2018-caleb-hannah-alex-1.jpg"},
				&Photo{Src: "thanksgiving-2018-caleb-hannah-alex-2.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-1.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-2.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-3.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-alex-daddy-grandy-1.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-alex-hannah-tickling-alex.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-alex-hannah-tickling-candler.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-alex.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-grandy-papa-1.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-grandy-papa-2.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-papa-ted-1.jpg"},
				&Photo{Src: "thanksgiving-2018-candler-papa-ted-2.jpg"},
				&Photo{Src: "thanksgiving-2018-dan-brandy-1.jpg"},
				&Photo{Src: "thanksgiving-2018-dan-brandy-2.jpg"},
				&Photo{Src: "thanksgiving-2018-frank-papa-bev-1.jpg"},
				&Photo{Src: "thanksgiving-2018-grandy-bev-1.jpg"},
				&Photo{Src: "thanksgiving-2018-grandy-bev-2.jpg"},
				&Photo{Src: "thanksgiving-2018-grandy-candler-alex.jpg"},
				&Photo{Src: "thanksgiving-2018-joseph-alex-1.jpg"},
				&Photo{Src: "thanksgiving-2018-mom-brandy.jpg"},
				&Photo{Src: "thanksgiving-2018-papa-ted-1.jpg"},
				&Photo{Src: "thanksgiving-2018-sam-hannah-1.jpg"},
				&Photo{Src: "thanksgiving-2018-sam-jennifer-caleb-hannah-1.jpg"},
				&Photo{Src: "thanksgiving-2018-sam-jennifer-caleb-hannah-2.jpg"},
				&Photo{Src: "thanksgiving-2018-sam-jennifer.jpg"},
			},
		},
		&Album{
			Title:       "Candler Playing Around",
			Description: "",
			Key:         "candler",
			CoverPhoto:  "candler-bathtub-trains.jpg",
			Photos: []*Photo{
				&Photo{Src: "candler-orange-at-grandys.jpg"},
				&Photo{Src: "candler-umbrella-deer.jpg"},
				&Photo{Src: "candler-and-josh.jpg"},
				&Photo{Src: "candler-bridge-fist-up.jpg"},
				&Photo{Src: "candler-bridge-pose-1.jpg"},
				&Photo{Src: "candler-bridge-pose-2.jpg"},
				&Photo{Src: "candler-bridge-pose-3.jpg"},
				&Photo{Src: "brownie-monster.jpg"},
				&Photo{Src: "candler-bathtub-trains.jpg"},
				&Photo{Src: "candler-glasses-laughing.jpg"},
				&Photo{Src: "candler-piano-standing.jpg"},
				&Photo{Src: "candler-piano-sitting-1.jpg"},
				&Photo{Src: "candler-piano-snoopy.jpg"},
				&Photo{Src: "candler-playing-trains-1.jpg"},
				&Photo{Src: "candler-trains-knex-bridge.jpg"},
				&Photo{Src: "candler-trains-on-floor.jpg"},
				&Photo{Src: "candler-train-sylvester-1.jpg"},
				&Photo{Src: "candler-train-sylvester-2.jpg"},
				&Photo{Src: "candler-golf-range.jpg"},
				&Photo{Src: "candler-playing-trains-1.jpg"},
				&Photo{Src: "candler-snoopy-breakfast-casserole.jpg"},
			},
		},
		&Album{
			Title:       "Grandy and Papa!",
			Description: "",
			Key:         "grandy-papa",
			CoverPhoto:  "grandy-and-papa.jpg",
			Photos: []*Photo{
				&Photo{Src: "candler-grandy-blowing-bubbles.jpg"},
				&Photo{Src: "candler-grandy-train-marietta.jpg"},
				&Photo{Src: "candler-grandy-cracker-barrel-1.jpg"},
				&Photo{Src: "candler-grandy-cracker-barrel-2.jpg"},
				&Photo{Src: "candler-grandy-swing.jpg"},
				&Photo{Src: "candler-papa-train-agrirama.jpg"},
				&Photo{Src: "candler-papa-train-marietta.jpg"},
				&Photo{Src: "grandy-and-papa.jpg"},
				&Photo{Src: "candler-grandy-papas-house.jpg"},
				&Photo{Src: "candler-papa-playing.jpg"},
				&Photo{Src: "christmas-2017-candler-grandly-papa.jpg"},
			},
		},
		&Album{
			Title:       "With Amy & Family",
			Description: "",
			Key:         "amy-and-family",
			CoverPhoto:  "eva-tina-ngan-thanh.jpg",
			Photos: []*Photo{
				&Photo{Src: "amy-xinh-dep-1.jpg"},
				&Photo{Src: "amy-xinh-dep-2.jpg"},
				&Photo{Src: "candler-and-ngan.jpg"},
				&Photo{Src: "candler-ngan-thanh-ice-cream.jpg"},
				&Photo{Src: "eva-ngan-amy-candler-uncle.jpg"},
				&Photo{Src: "candler-ngan-thanh-pool.jpg"},
				&Photo{Src: "eva-and-uncle.jpg"},
				&Photo{Src: "eva-tina-ngan-thanh.jpg"},
				&Photo{Src: "eva-and-grandfather-1.jpg"},
				&Photo{Src: "eva-and-grandfather-2.jpg"},
			},
		},
		&Album{
			Title:       "Random Pictures",
			Description: "",
			Key:         "random",
			CoverPhoto:  "eclipse-2017-2.jpg",
			Photos: []*Photo{
				&Photo{Description: "Papa Ted @ Ed's Truckstop",
					Src: "papa-ted-eds-truck-stop.jpg"},
				&Photo{Description: "Near Marietta Square",
					Src: "train-marietta.jpg"},
				&Photo{Description: "Eclipse 2017",
					Src: "eclipse-2017-1.jpg"},
				&Photo{Src: "eclipse-2017-2.jpg"},
			},
		},
		&Album{
			Title:       "Christmas 2016",
			Description: "",
			Key:         "christmas-2016",
			CoverPhoto:  "christmas-2016-75.jpg",
			Photos: func() []*Photo {
				ps := []*Photo{}
				for i := 0; i < 63; i++ {
					s := fmt.Sprintf("christmas-2016-%d.jpg", i)
					ps = append(ps, &Photo{Src: s})
				}
				return ps
			}(),
		},
		&Album{
			Title:       "Thanksgiving & Random 2016",
			Description: "",
			Key:         "thanksgiving-2016",
			CoverPhoto:  "thanksgiving-2016-55.jpg",
			Photos: func() []*Photo {
				ps := []*Photo{}
				is := []int{
					2, 4, 5, 10, 13, 14, 15, 18, 19, 20,
					22, 23, 25, 26, 30, 31, 33, 35, 37,
					43, 44, 49, 50, 55,
				}
				for _, n := range is {
					s := fmt.Sprintf("thanksgiving-2016-%d.jpg", n)
					ps = append(ps, &Photo{Src: s})
				}
				return ps
			}(),
		},
	}

	funcLog.WithMessage("about to insert").Info()

	// insert the albums
	for _, album := range albums {

		funcLog.WithMessagef("begin create album (key '%s')", album.Key).Info()

		// insert album and return if error
		album, err = CreatePhotoAlbum(txo, album)
		if err = txo.RollbackOnError(err); err != nil {
			return apierr.Errorf(err, funcTag, "%+v", album)
		}

		funcLog.WithMessage("done create album").Info()

		// insert the photos
		// could do a bulk insert here, but what the heck?
		for _, photo := range album.Photos {

			photo.AlbumID = album.ID
			photo.AlbumKey = album.Key

			funcLog.WithMessagef("begin create photo (album '%s', src '%s')", photo.AlbumKey, photo.Src).Info()

			// insert photo and return if error
			photo, err := CreatePhoto(txo, photo)
			if err = txo.RollbackOnError(err); err != nil {
				return apierr.Errorf(err, funcTag, "%+v", photo)
			}

			funcLog.WithMessage("done create photo").Info()
		}

	}

	return nil
}
