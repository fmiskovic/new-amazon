package repositories

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
)

func (s *RepositoryTestSuite) TestGetItemById() {
	repo := NewItemRepository(s.testDb.BunDb)

	s.Run("should return item by id", func() {
		// given
		itemId := uuid.MustParse("200cea28-b2b0-4051-9eb6-9a99e451af01")
		// when
		item, err := repo.GetById(s.testDb.Ctx, itemId)
		// then
		s.Nil(err)
		s.Equal("Cool Book 1", item.Title)
	})

	s.Run("should return error if item not found", func() {
		// given
		itemId := uuid.MustParse("200cea28-b2b0-4051-9eb6-9a99e451af10")
		// when
		_, err := repo.GetById(s.testDb.Ctx, itemId)
		// then
		s.NotNil(err)
	})
}

func (s *RepositoryTestSuite) TestGetItemsPage() {
	repo := NewItemRepository(s.testDb.BunDb)

	s.Run("should return page of items sorted by title in desc order", func() {
		// given
		pageRequest := entities.Pageable{
			Size:   2,
			Offset: 0,
			Sort: entities.Sort{
				Orders: []*entities.SortOrder{
					{
						Property:  "title",
						Direction: entities.DESC,
					},
				},
			},
		}
		// when
		page, err := repo.GetPage(s.testDb.Ctx, pageRequest)
		// then
		s.Nil(err)
		s.Len(page.Elements, 2)
		s.Equal(5, page.TotalElements)
		s.Equal("Cool Book 5", page.Elements[0].Title)
	})

	s.Run("given zero page request should return page of items", func() {
		// given
		pageRequest := entities.Pageable{}
		// when
		page, err := repo.GetPage(s.testDb.Ctx, pageRequest)
		// then
		s.Nil(err)
		s.Len(page.Elements, 5)
		s.Equal(5, page.TotalElements)
	})
}
