package operators

// Update removes then installs a
// package, to ensure it is up to date
func Update(pkg string) error {
	err := Remove(pkg)
	if err != nil {
		return err
	}

	err = Install(pkg)
	if err != nil {
		return err
	}

	return nil
}
