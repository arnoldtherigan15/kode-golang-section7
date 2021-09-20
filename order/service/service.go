package service

import (
	"errors"
	"kode-golang-section7/domain"
)

type service struct {
	repository domain.OrderRepository
}

func NewService(repository domain.OrderRepository) *service {
	return &service{repository}
}

func (s *service) Create(car *domain.Order) (*domain.Order, error) {
	createdCar, err := s.repository.Create(car)
	if err != nil {
		return &domain.Order{}, err
	}
	return createdCar, nil
}

func (s *service) FindAll() ([]*domain.Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) Update(order *domain.Order) (*domain.Order, error) {
	updatedCar, err := s.repository.Update(order)
	if err != nil {
		return updatedCar, err
	}
	return updatedCar, nil
}

func (s *service) Delete(ID int) (bool, error) {
	order, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if order.OrderID == 0 {
		return false, errors.New("order not found")
	}
	isDeleted, err := s.repository.Delete(order)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
